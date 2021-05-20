package main

import "fmt"

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