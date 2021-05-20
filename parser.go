package main

import (
	"fmt"
	"strconv"
)

func parse(tokens []Token) (count int, nodes []Node) {
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		if token.tokenType == TOKEN_IDENTIFIER && isKeyword(token.value) {
			// This is a statement
			switch string(token.value) {
			case "If":
				i++

				var exprTokens []Token
				for ; string(tokens[i].value) != "Then"; i++ {
					exprTokens = append(exprTokens, tokens[i])
				}
				i++
				exprTokens = append(exprTokens, Token{TOKEN_SEPARATOR, []byte(";")})
				count, cond := parse(exprTokens)

				var ifBody []Token
				level := 0
				for ; (string(tokens[i].value) != "End" && string(tokens[i].value) != "Else") || level != 0; i++ {
					if string(tokens[i].value) == "Func" ||
						string(tokens[i].value) == "Class" ||
						string(tokens[i].value) == "Singleton" ||
						string(tokens[i].value) == "If" ||
						string(tokens[i].value) == "While" ||
						string(tokens[i].value) == "For" {
						level++
					} else if string(tokens[i].value) == "End" {
						level--
					}
					ifBody = append(ifBody, tokens[i])
				}
				_, ifBodyNodes := parse(ifBody)
				if string(tokens[i].value) == "Else" {
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

				name := string(tokens[i].value) // TODO: Assert the name is an identifier
				i += 2                          // TODO: Assert the next token is an opening bracket
				var params []string
				for ; string(tokens[i].value) != ")"; i++ {
					params = append(params, string(tokens[i].value))
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
				if string(tokens[i+1].value) == "," {
					for ; string(tokens[i+1].value) == ","; i += 2 {
						names = append(names, string(tokens[i].value))
					}
					names = append(names, string(tokens[i].value))
				} else {
					names = append(names, string(tokens[i].value))
				}
				nodes = append(nodes, PropNode{names})
				count++
			case "For":
				i += 2

				// Expressions within the brackets
				var initTokens []Token
				for ; string(tokens[i].value) != ";"; i++ {
					initTokens = append(initTokens, tokens[i])
				}
				initTokens = append(initTokens, tokens[i])
				i++
				_, init := parse(initTokens)
				var condTokens []Token
				for ; string(tokens[i].value) != ";"; i++ {
					condTokens = append(condTokens, tokens[i])
				}
				condTokens = append(condTokens, tokens[i])
				i++
				_, cond := parse(condTokens)
				var incTokens []Token
				for ; string(tokens[i].value) != ")"; i++ {
					incTokens = append(incTokens, tokens[i])
				}
				incTokens = append(incTokens, Token{TOKEN_SEPARATOR, []byte(";")})
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
				for ; string(tokens[i].value) != "Then"; i++ {
					exprTokens = append(exprTokens, tokens[i])
				}
				i++
				exprTokens = append(exprTokens, Token{TOKEN_SEPARATOR, []byte(";")})
				_, cond := parse(exprTokens)

				var bodyTokens []Token
				takeUntilEnd(tokens, &bodyTokens, &i)
				_, body := parse(bodyTokens)

				nodes = append(nodes, WhileNode{cond[0], body})
				count++
			case "Class", "Singleton":
				isSingleton := string(tokens[i].value) == "Singleton"
				i++
				name := string(tokens[i].value)
				var extends []string
				if i++; string(tokens[i].value) == "Extends" {
					i++
					if string(tokens[i+1].value) == "," {
						for ; string(tokens[i+1].value) == ","; i += 2 {
							extends = append(extends, string(tokens[i].value))
						}
					} else {
						extends = append(extends, string(tokens[i].value))
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
				for ; string(tokens[i].value) != ";"; i++ {
					expressionTokens = append(expressionTokens, tokens[i])
				}
				i++
				_, expression := parse(expressionTokens)
				nodes = append(nodes, ReturnNode{expression})
				count++
			}
		} else {
			// This is an expression
			savedIndex := i
			hasOperator := false
			level := 0
			for ; i < len(tokens) && string(tokens[i].value) != ";"; i++ {
				if string(tokens[i].value) == "(" || string(tokens[i].value) == "[" {
					level++
				}
				if string(tokens[i].value) == ")" || string(tokens[i].value) == "]" {
					level--
				}
				if tokens[i].tokenType == TOKEN_OPERATOR && level == 0 {
					hasOperator = true
				}
			}
			i = savedIndex
			if hasOperator {
				// Resolve expression
				savedIndex := i
			presendenceLoop:
				for _, o := range operators {
					for ; string(tokens[i].value) != ";"; i++ {
						if string(tokens[i].value) == o {
							_, LHS := parse(tokens[savedIndex:i])
							savedIndex = i
							for string(tokens[i].value) != ";" {
								i++
							}
							_, RHS := parse(tokens[savedIndex+1 : i+1])
							nodes = append(nodes, BinopNode{o, LHS[0], RHS[0]})
							count++
							break presendenceLoop
						}
					}
					i = savedIndex
				}

			} else {
				// Single non-statement node
				if tokens[i].tokenType == TOKEN_NUMBER {
					number, _ := strconv.Atoi(string(tokens[i].value))
					nodes = append(nodes, NumberNode{number})
					count++
					i++
					// }
				} else if string(token.value) == "[" {
					level := 0
					i++
					var arrayContentNodes []Node
					for {
						var elementTokens []Token
						for {
							if string(tokens[i].value) == "[" {
								level++
							} else if string(tokens[i].value) == "]" {
								level--
							}

							if (string(tokens[i].value) == "]" && level < 0) || (string(tokens[i].value) == "," && level == 0) {
								if string(tokens[i].value) == "]" {
									elementTokens = append(elementTokens, tokens[i])
								}
								break
							}
							elementTokens = append(elementTokens, tokens[i])
							i++
						}
						elementTokens = append(elementTokens, Token{TOKEN_SEPARATOR, []byte(";")})
						_, element := parse(elementTokens)
						arrayContentNodes = append(arrayContentNodes, element[0])
						if string(tokens[i].value) == "]" && level < 0 {
							break
						}
						i++
					}
					nodes = append(nodes, ArrayNode{arrayContentNodes})
					count = 1
					return
				} else if tokens[i].tokenType == TOKEN_STRINGLITERAL {
					nodes = append(nodes, StringLiteralNode{string(tokens[i].value)})
					count++
				} else if len(tokens) > 1 && tokens[i].tokenType == TOKEN_IDENTIFIER && string(tokens[i+1].value) == "(" {
					// Function call
					name := string(tokens[i].value)
					i += 2
					var args []Node
					level := 0
					for {
						var elementTokens []Token
						for {
							if string(tokens[i].value) == "(" {
								level++
							}
							if string(tokens[i].value) == ")" {
								level--
							}

							if (string(tokens[i].value) == ")" && level < 0) ||
								(string(tokens[i].value) == "," && level == 0) {
								break
							}

							elementTokens = append(elementTokens, tokens[i])
							i++
						}
						elementTokens = append(elementTokens, Token{TOKEN_SEPARATOR, []byte(";")})
						_, arg := parse(elementTokens)
						args = append(args, arg[0])
						if string(tokens[i].value) == ")" && level < 0 {
							break
						}
						i++
					}
					nodes = append(nodes, FunctionCallNode{name, args})
					count++
				} else if tokens[i].tokenType == TOKEN_IDENTIFIER {
					nodes = append(nodes, IdentifierNode{string(tokens[i].value)})
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
	for ; string(tokens[*i].value) != "End" || level != 0; (*i)++ {
		if string(tokens[*i].value) == "Func" ||
			string(tokens[*i].value) == "Class" ||
			string(tokens[*i].value) == "Singleton" ||
			string(tokens[*i].value) == "If" ||
			string(tokens[*i].value) == "While" ||
			string(tokens[*i].value) == "For" {
			level++
		} else if string(tokens[*i].value) == "End" {
			level--
		}
		*output = append(*output, tokens[*i])
	}
}

var operators = [...]string{"=", "||", "&&", "|", "&", "==", "!=", ">", ">=", "<", "<=", "<<", ">>", "+", "-", "*", "/", "!", "~", "."}
var keywords = [...]string{"If", "While", "Class", "End", "Singleton", "For", "Func", "Return", "Else", "Extends", "Prop"}

func isKeyword(identifier []byte) bool {
	for _, k := range keywords {
		if k == string(identifier) {
			return true
		}
	}
	return false
}

// Types

const (
	NODE_IF            = iota
	NODE_FUNC          = iota
	NODE_PROP          = iota
	NODE_CLASS         = iota
	NODE_RETURN        = iota
	NODE_BINOP         = iota
	NODE_IDENTIFIER    = iota
	NODE_NUMBER        = iota
	NODE_FOR           = iota
	NODE_WHILE         = iota
	NODE_ARRAY         = iota
	NODE_FUNCTIONCALL  = iota
	NODE_STRINGLITERAL = iota
)

type Node interface {
	getType() int
	String() string
}

type IfNode struct {
	condition Node
	ifbody    []Node
	elsebody  []Node
}

func (n IfNode) getType() int {
	return NODE_IF
}

func (n IfNode) String() string {
	return fmt.Sprintf("If node: {\n cond: %s,\n ifBody: %s\n elseBody: %s\n}\n", n.condition, n.ifbody, n.elsebody)
}

type FuncNode struct {
	name   string
	params []string
	body   []Node
}

func (n FuncNode) getType() int {
	return NODE_FUNC
}

func (n FuncNode) String() string {
	return fmt.Sprintf("Function node: {\n name: %s,\n params: %s\n body: %s\n}\n", n.name, n.params, n.body)
}

type PropNode struct {
	names []string
}

func (n PropNode) getType() int {
	return NODE_PROP
}

func (n PropNode) String() string {
	return fmt.Sprintf("Prop node: {\n names: %s\n}\n", n.names)
}

type ClassNode struct {
	isSingleton bool
	name        string
	extends     []string
	body        []Node
}

func (n ClassNode) getType() int {
	return NODE_CLASS
}
func (n ClassNode) String() string {
	return fmt.Sprintf("Class node: {\n singleton: %t,\n name: %s,\n extends: %s\n body: %s\n}\n", n.isSingleton, n.name, n.extends, n.body)
}

type ReturnNode struct {
	expression []Node
}

func (n ReturnNode) getType() int {
	return NODE_RETURN
}
func (n ReturnNode) String() string {
	return fmt.Sprintf("Return node: {\n expression: %s\n}\n", n.expression)
}

type BinopNode struct {
	binopType string
	LHS       Node
	RHS       Node
}

func (n BinopNode) getType() int {
	return NODE_BINOP
}
func (n BinopNode) String() string {
	return fmt.Sprintf("Binop node: {\n type: %s,\n LHS: %s,\n RHS: %s\n}\n", n.binopType, n.LHS, n.RHS)
}

type IdentifierNode struct {
	name string
}

func (n IdentifierNode) getType() int {
	return NODE_IDENTIFIER
}
func (n IdentifierNode) String() string {
	return fmt.Sprintf("Identifier node: {\n name: %s\n}\n", n.name)
}

type NumberNode struct {
	value int
}

func (n NumberNode) getType() int {
	return NODE_NUMBER
}
func (n NumberNode) String() string {
	return fmt.Sprintf("Number node: {\n value: %d\n}\n", n.value)
}

type ForNode struct {
	init Node
	cond Node
	inc  Node
	body []Node
}

func (n ForNode) getType() int {
	return NODE_FOR
}
func (n ForNode) String() string {
	return fmt.Sprintf("For node: {\n init: %s,\n cond: %s,\n inc: %s,\n body: %s\n}\n", n.init, n.cond, n.inc, n.body)
}

type WhileNode struct {
	cond Node
	body []Node
}

func (n WhileNode) getType() int {
	return NODE_WHILE
}
func (n WhileNode) String() string {
	return fmt.Sprintf("While node: {\n cond: %s,\n body: %s\n}\n", n.cond, n.body)
}

type ArrayNode struct {
	items []Node
}

func (n ArrayNode) getType() int {
	return NODE_ARRAY
}
func (n ArrayNode) String() string {
	return fmt.Sprintf("Array node: {\n items: %s\n}\n", n.items)
}

type FunctionCallNode struct {
	name string
	args []Node
}

func (n FunctionCallNode) getType() int {
	return NODE_FUNCTIONCALL
}
func (n FunctionCallNode) String() string {
	return fmt.Sprintf("Function call node: {\n name: %s,\n args: %s\n}\n", n.name, n.args)
}

type StringLiteralNode struct {
	literal string
}

func (n StringLiteralNode) getType() int {
	return NODE_STRINGLITERAL
}
func (n StringLiteralNode) String() string {
	return fmt.Sprintf("String literal node: {\n literal: %s\n}\n", n.literal)
}
