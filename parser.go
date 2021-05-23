package main

import "strconv"

func parse(tokens []Token) (count int, nodes []Node) {
	var prevNode Node
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		if token.tokenType == TOKEN_IDENTIFIER && isKeyword(token.value) {
			// This is a statement
			switch token.value {
			case "If":
				i++

				var exprTokens []Token
				for ; tokens[i].value != "Then"; i++ {
					exprTokens = append(exprTokens, tokens[i])
				}
				i++
				exprTokens = append(exprTokens, Token{TOKEN_SEPARATOR, ";"})
				count, cond := parse(exprTokens)

				var ifBody []Token
				level := 0
				for ; (tokens[i].value != "End" && tokens[i].value != "Else") || level != 0; i++ {
					if tokens[i].value == "Func" ||
						tokens[i].value == "Class" ||
						tokens[i].value == "Singleton" ||
						tokens[i].value == "If" ||
						tokens[i].value == "While" ||
						tokens[i].value == "For" {
						level++
					} else if tokens[i].value == "End" {
						level--
					}
					ifBody = append(ifBody, tokens[i])
				}
				_, ifBodyNodes := parse(ifBody)
				if tokens[i].value == "Else" {
					var elseBlock []Token
					takeUntilEnd(tokens, &elseBlock, &i)
					_, elseBlockNodes := parse(elseBlock)
					prevNode = IfNode{cond[0], ifBodyNodes, elseBlockNodes}
					nodes = append(nodes, prevNode)
				} else {
					prevNode = IfNode{cond[0], ifBodyNodes, nil}
					nodes = append(nodes, prevNode)
				}

				count++
			case "Func":
				i++

				name := tokens[i].value // TODO: Assert the name is an identifier
				i += 2                  // TODO: Assert the next token is an opening bracket
				var params []string
				for ; tokens[i].value != ")"; i++ {
					if tokens[i].value != "," { // FIXME
						params = append(params, tokens[i].value)
					}
				}
				i++
				var bodyTokens []Token
				takeUntilEnd(tokens, &bodyTokens, &i)
				_, body := parse(bodyTokens)

				prevNode = FuncNode{name, params, body}
				nodes = append(nodes, prevNode)
				count++
			case "For":
				i += 2

				// Expressions within the brackets
				var initTokens []Token
				for ; tokens[i].value != ";"; i++ {
					initTokens = append(initTokens, tokens[i])
				}
				initTokens = append(initTokens, tokens[i])
				i++
				_, init := parse(initTokens)
				var condTokens []Token
				for ; tokens[i].value != ";"; i++ {
					condTokens = append(condTokens, tokens[i])
				}
				condTokens = append(condTokens, tokens[i])
				i++
				_, cond := parse(condTokens)
				var incTokens []Token
				for ; tokens[i].value != ")"; i++ {
					incTokens = append(incTokens, tokens[i])
				}
				incTokens = append(incTokens, Token{TOKEN_SEPARATOR, ";"})
				i++
				_, inc := parse(incTokens)

				var bodyTokens []Token
				takeUntilEnd(tokens, &bodyTokens, &i)
				_, body := parse(bodyTokens)

				prevNode = ForNode{init[0], cond[0], inc[0], body}
				nodes = append(nodes, prevNode)
				count++
			case "While":
				i++
				var exprTokens []Token
				for ; tokens[i].value != "Then"; i++ {
					exprTokens = append(exprTokens, tokens[i])
				}
				i++
				exprTokens = append(exprTokens, Token{TOKEN_SEPARATOR, ";"})
				_, cond := parse(exprTokens)

				var bodyTokens []Token
				takeUntilEnd(tokens, &bodyTokens, &i)
				_, body := parse(bodyTokens)

				prevNode = WhileNode{cond[0], body}
				nodes = append(nodes, prevNode)
				count++
			case "Return":
				i++
				var expressionTokens []Token
				for ; tokens[i].value != ";"; i++ {
					expressionTokens = append(expressionTokens, tokens[i])
				}
				i++
				_, expression := parse(expressionTokens)
				prevNode = ReturnNode{expression[0]}
				nodes = append(nodes, prevNode)
				count++
			case "Fold":
				i++ // eat keyword
				name := tokens[i].value
				i++
				var fields []string
				for {
					if tokens[i].value == "End" {
						break
					}
					fields = append(fields, tokens[i].value)
					i++
					if tokens[i].value == "," {
						i++
						continue
					} else {
						break
					}
				}
				prevNode = FoldNode{name, fields}
				nodes = append(nodes, prevNode)
				count++
			}
		} else {
			// This is an expression
			savedIndex := i
			hasOperator := false
			minLevel := 1000
			level := 0
			for ; i < len(tokens) && tokens[i].value != ";"; i++ {
				if tokens[i].value == "(" || tokens[i].value == "[" {
					level++
				}
				if tokens[i].value == ")" || tokens[i].value == "]" {
					level--
				}
				if tokens[i].tokenType == TOKEN_OPERATOR {
					hasOperator = true
					if minLevel > level {
						minLevel = level
					}
				}
			}
			i = savedIndex
			if hasOperator && (minLevel < 1 || (tokens[i].value == "(" && (prevNode == nil || (prevNode.getType() != NODE_IDENTIFIER && prevNode.getType() != NODE_ARRAY && prevNode.getType() != NODE_STRINGLITERAL && prevNode.getType() == NODE_FUNCTIONCALL)))) {
				// Resolve expression
				savedIndex := i
				level := 0
			presendenceLoop:
				for _, o := range operators {
					for ; i < len(tokens) && tokens[i].value != ";"; i++ {
						if tokens[i].value == "(" || tokens[i].value == "[" {
							level++
						} else if tokens[i].value == ")" || tokens[i].value == "]" {
							level--
						} else if tokens[i].value == o && level == minLevel {
							_, LHS := parse(tokens[savedIndex+minLevel : i])
							savedIndex = i
							for i < len(tokens) && tokens[i].value != ";" {
								i++
							}
							_, RHS := parse(tokens[savedIndex+1 : i-minLevel])
							prevNode = BinopNode{o, LHS[0], RHS[0]}
							nodes = append(nodes, prevNode)
							count++
							break presendenceLoop
						}
					}
					i = savedIndex
					level = 0
				}
			} else {
				// Single non-statement node
				if tokens[i].tokenType == TOKEN_NUMBER {
					number, _ := strconv.ParseFloat(tokens[i].value, 64)
					prevNode = NumberNode{number}
					nodes = append(nodes, prevNode)
					count++
					i++
				} else if token.value == "[" {
					level := 0
					i++
					var arrayContentNodes []Node
					for {
						var elementTokens []Token
						for {
							if tokens[i].value == "[" {
								level++
							} else if tokens[i].value == "]" {
								level--
							}

							if (tokens[i].value == "]" && level < 0) || (tokens[i].value == "," && level == 0) {
								// if tokens[i].value == "]" {
								// 	elementTokens = append(elementTokens, tokens[i])
								// }
								break
							}
							elementTokens = append(elementTokens, tokens[i])
							i++
						}
						elementTokens = append(elementTokens, Token{TOKEN_SEPARATOR, ";"})
						_, element := parse(elementTokens)
						arrayContentNodes = append(arrayContentNodes, element[0])
						if tokens[i].value == "]" && level < 0 {
							break
						}
						i++
					}
					prevNode = ArrayNode{arrayContentNodes}
					nodes = append(nodes, prevNode)
					count++
				} else if tokens[i].tokenType == TOKEN_STRINGLITERAL {
					prevNode = StringLiteralNode{tokens[i].value[1 : len(tokens[i].value)-1]}
					nodes = append(nodes, prevNode)
					count++
				} else if prevNode != nil && (prevNode.getType() == NODE_IDENTIFIER || prevNode.getType() == NODE_ARRAY || prevNode.getType() == NODE_STRINGLITERAL || prevNode.getType() == NODE_FUNCTIONCALL) && tokens[i].value == "(" {
					// Function call
					name := prevNode
					i++
					var args []Node
					level := 0
					if tokens[i].value != ")" {
						for {
							var elementTokens []Token
							for {
								if tokens[i].value == "(" || tokens[i].value == "[" {
									level++
								} else if tokens[i].value == ")" || tokens[i].value == "]" {
									level--
								}

								if (tokens[i].value == ")" && level < 0) ||
									(tokens[i].value == "," && level == 0) {
									break
								}

								elementTokens = append(elementTokens, tokens[i])
								i++
							}
							elementTokens = append(elementTokens, Token{TOKEN_SEPARATOR, ";"})
							_, arg := parse(elementTokens)
							args = append(args, arg[0])
							if tokens[i].value == ")" && level < 0 {
								break
							}
							i++
						}
						prevNode = FunctionCallNode{name, args}
						nodes[len(nodes)-1] = prevNode
						count++
					} else {
						prevNode = FunctionCallNode{name, nil}
						nodes[len(nodes)-1] = prevNode
						count++
					}
				} else if tokens[i].tokenType == TOKEN_IDENTIFIER {
					prevNode = IdentifierNode{tokens[i].value}
					nodes = append(nodes, prevNode)
					count++
				}
			}
		}
	}
	return
}

func takeUntilEnd(tokens []Token, output *[]Token, i *int) {
	level := 0
	for ; tokens[*i].value != "End" || level != 0; (*i)++ {
		if tokens[*i].value == "Func" ||
			tokens[*i].value == "Class" ||
			tokens[*i].value == "Singleton" ||
			tokens[*i].value == "If" ||
			tokens[*i].value == "While" ||
			tokens[*i].value == "For" {
			level++
		} else if tokens[*i].value == "End" {
			level--
		}
		*output = append(*output, tokens[*i])
	}
}

var operators = [...]string{"=", "||", "&&", "|", "&", "^", "==", "!=", ">", ">=", "<", "<=", "<<", ">>", "+", "-", "*", "/", "!", "~", "."}
var keywords = [...]string{"If", "While", "End", "For", "Func", "Return", "Else", "Fold"}

func isKeyword(identifier string) bool {
	for _, k := range keywords {
		if k == identifier {
			return true
		}
	}
	return false
}
