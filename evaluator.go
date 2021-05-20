package main

import "fmt"

var variables map[string]interface{}

type Function struct {
	params []string
	body   []Node
}

func (n BinopNode) evaluate() (r interface{}) {
	switch n.binopType {
	case "+":
		LHS := n.LHS.evaluate()
		RHS := n.RHS.evaluate()
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			r = n.LHS.evaluate().(int) + n.RHS.evaluate().(int)
		}
	case "-":
		LHS := n.LHS.evaluate()
		RHS := n.RHS.evaluate()
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			r = n.LHS.evaluate().(int) - n.RHS.evaluate().(int)
		}
	case "*":
		LHS := n.LHS.evaluate()
		RHS := n.RHS.evaluate()
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			r = n.LHS.evaluate().(int) * n.RHS.evaluate().(int)
		}
	case "/":
		LHS := n.LHS.evaluate()
		RHS := n.RHS.evaluate()
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			r = n.LHS.evaluate().(int) / n.RHS.evaluate().(int)
		}
	case "<<":
		LHS := n.LHS.evaluate()
		RHS := n.RHS.evaluate()
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			r = n.LHS.evaluate().(int) << n.RHS.evaluate().(int)
		}
	case ">>":
		LHS := n.LHS.evaluate()
		RHS := n.RHS.evaluate()
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			r = n.LHS.evaluate().(int) >> n.RHS.evaluate().(int)
		}
	case "&":
		LHS := n.LHS.evaluate()
		RHS := n.RHS.evaluate()
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			r = n.LHS.evaluate().(int) & n.RHS.evaluate().(int)
		}
	case "|":
		LHS := n.LHS.evaluate()
		RHS := n.RHS.evaluate()
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			r = n.LHS.evaluate().(int) | n.RHS.evaluate().(int)
		}
	case "^":
		LHS := n.LHS.evaluate()
		RHS := n.RHS.evaluate()
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			r = n.LHS.evaluate().(int) ^ n.RHS.evaluate().(int)
		}
	case "==":
		LHS := n.LHS.evaluate()
		RHS := n.RHS.evaluate()
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			if n.LHS.evaluate().(int) == n.RHS.evaluate().(int) {
				r = 1
			} else {
				r = 0
			}
		}
	case "!=":
		LHS := n.LHS.evaluate()
		RHS := n.RHS.evaluate()
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			if n.LHS.evaluate().(int) != n.RHS.evaluate().(int) {
				r = 1
			} else {
				r = 0
			}
		}
	case "<":
		LHS := n.LHS.evaluate()
		RHS := n.RHS.evaluate()
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			if n.LHS.evaluate().(int) < n.RHS.evaluate().(int) {
				r = 1
			} else {
				r = 0
			}
		}
	case "<=":
		LHS := n.LHS.evaluate()
		RHS := n.RHS.evaluate()
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			if n.LHS.evaluate().(int) <= n.RHS.evaluate().(int) {
				r = 1
			} else {
				r = 0
			}
		}
	case ">":
		LHS := n.LHS.evaluate()
		RHS := n.RHS.evaluate()
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			if n.LHS.evaluate().(int) > n.RHS.evaluate().(int) {
				r = 1
			} else {
				r = 0
			}
		}
	case ">=":
		LHS := n.LHS.evaluate()
		RHS := n.RHS.evaluate()
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			if n.LHS.evaluate().(int) >= n.RHS.evaluate().(int) {
				r = 1
			} else {
				r = 0
			}
		}
	case "&&":
		LHS := n.LHS.evaluate()
		RHS := n.RHS.evaluate()
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			if n.LHS.evaluate().(int) != 0 && n.RHS.evaluate().(int) != 0 {
				r = 1
			} else {
				r = 0
			}
		}
	case "||":
		LHS := n.LHS.evaluate()
		RHS := n.RHS.evaluate()
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			if n.LHS.evaluate().(int) != 0 || n.RHS.evaluate().(int) != 0 {
				r = 1
			} else {
				r = 0
			}
		}
	case "=":
		RHS := n.RHS.evaluate()
		_, RHSok := RHS.(int)
		if RHSok {
			variables[n.LHS.String()] = RHS
			r = RHS
		}
	}
	return
}

func (n FuncNode) evaluate() interface{} {
	variables[n.name] = Function{n.params, n.body}
	return nil
}
func (n ClassNode) evaluate() interface{} {
	return 0
}
func (n NumberNode) evaluate() interface{} {
	return n.value
}
func (n StringLiteralNode) evaluate() interface{} {
	return n.literal
}
func (n PropNode) evaluate() interface{} {
	return 0
}
func (n IfNode) evaluate() interface{} {
	if n.condition.evaluate().(int) != 0 {
		for _, b := range n.ifbody {
			fmt.Printf("%v\n", b.evaluate())
		}
	} else {
		if n.elsebody != nil {
			for _, b := range n.elsebody {
				fmt.Printf("%v\n", b.evaluate())
			}
		}
	}
	return nil
}
func (n ForNode) evaluate() interface{} {
	n.init.evaluate()
	for {
		if n.cond.evaluate().(int) == 0 {
			break
		}

		for _, b := range n.body {
			fmt.Printf("%v\n", b.evaluate())
		}

		n.inc.evaluate()
	}
	return nil
}
func (n WhileNode) evaluate() interface{} {
	for {
		if n.cond.evaluate().(int) == 0 {
			break
		}
		for _, b := range n.body {
			fmt.Printf("%v\n", b.evaluate())
		}
	}
	return nil
}
func (n ReturnNode) evaluate() interface{} {
	return 0
}
func (n ArrayNode) evaluate() interface{} {
	return 0
}
func (n FunctionCallNode) evaluate() interface{} {
	function := variables[n.name]
	for _, b := range function.(Function).body {
		b.evaluate()
	}
	return nil
}
func (n IdentifierNode) evaluate() interface{} {
	return variables[n.name]
}
