package main

import "strconv"

func parse(tokens []Token) (count int, nodes []Node) {
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
					nodes = append(nodes, IfNode{cond[0], ifBodyNodes, elseBlockNodes})
				} else {
					nodes = append(nodes, IfNode{cond[0], ifBodyNodes, nil})
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

				nodes = append(nodes, FuncNode{name, params, body})
				count++
			case "Prop":
				i++
				var names []string
				if tokens[i+1].value == "," {
					for ; tokens[i+1].value == ","; i += 2 {
						names = append(names, tokens[i].value)
					}
					names = append(names, tokens[i].value)
				} else {
					names = append(names, tokens[i].value)
				}
				nodes = append(nodes, PropNode{names})
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

				nodes = append(nodes, ForNode{init[0], cond[0], inc[0], body})
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

				nodes = append(nodes, WhileNode{cond[0], body})
				count++
			case "Class", "Singleton":
				isSingleton := tokens[i].value == "Singleton"
				i++
				name := tokens[i].value
				var extends []string
				if i++; tokens[i].value == "Extends" {
					i++
					if tokens[i+1].value == "," {
						for ; tokens[i+1].value == ","; i += 2 {
							extends = append(extends, tokens[i].value)
						}
					} else {
						extends = append(extends, tokens[i].value)
						i++
					}
				}
				var bodyTokens []Token
				takeUntilEnd(tokens, &bodyTokens, &i)
				_, body := parse(bodyTokens)
				nodes = append(nodes, ClassNode{isSingleton, name, extends, body})
				count++
			case "Return":
				i++
				var expressionTokens []Token
				for ; tokens[i].value != ";"; i++ {
					expressionTokens = append(expressionTokens, tokens[i])
				}
				i++
				_, expression := parse(expressionTokens)
				nodes = append(nodes, ReturnNode{expression[0]})
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
			if hasOperator {
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
							nodes = append(nodes, BinopNode{o, LHS[0], RHS[0]})
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
					number, _ := strconv.Atoi(tokens[i].value)
					nodes = append(nodes, NumberNode{number})
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
								if tokens[i].value == "]" {
									elementTokens = append(elementTokens, tokens[i])
								}
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
					nodes = append(nodes, ArrayNode{arrayContentNodes})
					count = 1
					return
				} else if tokens[i].tokenType == TOKEN_STRINGLITERAL {
					nodes = append(nodes, StringLiteralNode{tokens[i].value})
					count++
				} else if len(tokens) > 1 && tokens[i].tokenType == TOKEN_IDENTIFIER && tokens[i+1].value == "(" {
					// Function call
					name := tokens[i].value
					i += 2
					var args []Node
					level := 0
					if tokens[i].value != ")" {
						for {
							var elementTokens []Token
							for {
								if tokens[i].value == "(" {
									level++
								}
								if tokens[i].value == ")" {
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
						nodes = append(nodes, FunctionCallNode{name, args})
						count++
					} else {
						nodes = append(nodes, FunctionCallNode{name, nil})
						count++
					}
				} else if tokens[i].tokenType == TOKEN_IDENTIFIER {
					nodes = append(nodes, IdentifierNode{tokens[i].value})
					count++
					i++
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
var keywords = [...]string{"If", "While", "Class", "End", "Singleton", "For", "Func", "Return", "Else", "Extends", "Prop"}

func isKeyword(identifier string) bool {
	for _, k := range keywords {
		if k == identifier {
			return true
		}
	}
	return false
}
