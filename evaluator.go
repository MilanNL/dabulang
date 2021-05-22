package main

type Scope struct {
	parent    *Scope
	variables map[string]interface{}
}

func (s *Scope) addToScope(name string, value interface{}) {
	s.variables[name] = value
}

func (s *Scope) getFromScope(name string) interface{} {
	// b := &s
	// for {
	// 	if val, found := (*b).variables[name]; found {
	// 		return val
	// 	}
	// 	b = s.parent
	// }
	return s.variables[name]
}

func createScope(parent *Scope) Scope {
	return Scope{parent, make(map[string]interface{})}
}

var global_scope Scope = Scope{nil, make(map[string]interface{})}

type Function struct {
	params []string
	body   []Node
}

type Fold map[string]interface{}

var exit_function = false
var return_value interface{}

func (n BinopNode) evaluate(scope *Scope) (r interface{}) {
	switch n.binopType {
	case "+":
		LHS := n.LHS.evaluate(scope)
		RHS := n.RHS.evaluate(scope)
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			r = n.LHS.evaluate(scope).(int) + n.RHS.evaluate(scope).(int)
		}
	case "-":
		LHS := n.LHS.evaluate(scope)
		RHS := n.RHS.evaluate(scope)
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			r = n.LHS.evaluate(scope).(int) - n.RHS.evaluate(scope).(int)
		}
	case "*":
		LHS := n.LHS.evaluate(scope)
		RHS := n.RHS.evaluate(scope)
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			r = n.LHS.evaluate(scope).(int) * n.RHS.evaluate(scope).(int)
		}
	case "/":
		LHS := n.LHS.evaluate(scope)
		RHS := n.RHS.evaluate(scope)
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			r = n.LHS.evaluate(scope).(int) / n.RHS.evaluate(scope).(int)
		}
	case "<<":
		LHS := n.LHS.evaluate(scope)
		RHS := n.RHS.evaluate(scope)
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			r = n.LHS.evaluate(scope).(int) << n.RHS.evaluate(scope).(int)
		}
	case ">>":
		LHS := n.LHS.evaluate(scope)
		RHS := n.RHS.evaluate(scope)
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			r = n.LHS.evaluate(scope).(int) >> n.RHS.evaluate(scope).(int)
		}
	case "&":
		LHS := n.LHS.evaluate(scope)
		RHS := n.RHS.evaluate(scope)
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			r = n.LHS.evaluate(scope).(int) & n.RHS.evaluate(scope).(int)
		}
	case "|":
		LHS := n.LHS.evaluate(scope)
		RHS := n.RHS.evaluate(scope)
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			r = n.LHS.evaluate(scope).(int) | n.RHS.evaluate(scope).(int)
		}
	case "^":
		LHS := n.LHS.evaluate(scope)
		RHS := n.RHS.evaluate(scope)
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			r = n.LHS.evaluate(scope).(int) ^ n.RHS.evaluate(scope).(int)
		}
	case "==":
		LHS := n.LHS.evaluate(scope)
		RHS := n.RHS.evaluate(scope)
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			if n.LHS.evaluate(scope).(int) == n.RHS.evaluate(scope).(int) {
				r = 1
			} else {
				r = 0
			}
		}
	case "!=":
		LHS := n.LHS.evaluate(scope)
		RHS := n.RHS.evaluate(scope)
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			if n.LHS.evaluate(scope).(int) != n.RHS.evaluate(scope).(int) {
				r = 1
			} else {
				r = 0
			}
		}
	case "<":
		LHS := n.LHS.evaluate(scope)
		RHS := n.RHS.evaluate(scope)
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			if n.LHS.evaluate(scope).(int) < n.RHS.evaluate(scope).(int) {
				r = 1
			} else {
				r = 0
			}
		}
	case "<=":
		LHS := n.LHS.evaluate(scope)
		RHS := n.RHS.evaluate(scope)
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			if n.LHS.evaluate(scope).(int) <= n.RHS.evaluate(scope).(int) {
				r = 1
			} else {
				r = 0
			}
		}
	case ">":
		LHS := n.LHS.evaluate(scope)
		RHS := n.RHS.evaluate(scope)
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			if n.LHS.evaluate(scope).(int) > n.RHS.evaluate(scope).(int) {
				r = 1
			} else {
				r = 0
			}
		}
	case ">=":
		LHS := n.LHS.evaluate(scope)
		RHS := n.RHS.evaluate(scope)
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			if n.LHS.evaluate(scope).(int) >= n.RHS.evaluate(scope).(int) {
				r = 1
			} else {
				r = 0
			}
		}
	case "&&":
		LHS := n.LHS.evaluate(scope)
		RHS := n.RHS.evaluate(scope)
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			if n.LHS.evaluate(scope).(int) != 0 && n.RHS.evaluate(scope).(int) != 0 {
				r = 1
			} else {
				r = 0
			}
		}
	case "||":
		LHS := n.LHS.evaluate(scope)
		RHS := n.RHS.evaluate(scope)
		_, LHSok := LHS.(int)
		_, RHSok := RHS.(int)
		if LHSok && RHSok {
			if n.LHS.evaluate(scope).(int) != 0 || n.RHS.evaluate(scope).(int) != 0 {
				r = 1
			} else {
				r = 0
			}
		}
	case "=":
		RHS := n.RHS.evaluate(scope)
		scope.addToScope(n.LHS.String(), RHS)
		r = RHS
	}
	return
}

func (n FuncNode) evaluate(scope *Scope) interface{} {
	scope.addToScope(n.name, Function{n.params, n.body})
	return nil
}

func (n NumberNode) evaluate(_ *Scope) interface{} {
	return n.value
}
func (n StringLiteralNode) evaluate(_ *Scope) interface{} {
	return n.literal
}

func (n IfNode) evaluate(scope *Scope) interface{} {
	if n.condition.evaluate(scope).(int) != 0 {
		for _, b := range n.ifbody {
			b.evaluate(scope)
			if exit_function {
				return nil
			}
		}
	} else {
		if n.elsebody != nil {
			for _, b := range n.elsebody {
				b.evaluate(scope)
				if exit_function {
					return nil
				}
			}
		}
	}
	return nil
}

func (n ForNode) evaluate(scope *Scope) interface{} {
	n.init.evaluate(scope)
	for {
		if n.cond.evaluate(scope).(int) == 0 {
			break
		}

		for _, b := range n.body {
			b.evaluate(scope)
			if exit_function {
				return nil
			}
		}

		n.inc.evaluate(scope)
	}
	return nil
}

func (n WhileNode) evaluate(scope *Scope) interface{} {
	for {
		if n.cond.evaluate(scope).(int) == 0 {
			break
		}
		for _, b := range n.body {
			b.evaluate(scope)
			if exit_function {
				return nil
			}
		}
	}
	return nil
}

func (n ReturnNode) evaluate(scope *Scope) interface{} {
	exit_function = true
	if n.expression != nil {
		return_value = n.expression.evaluate(scope)
	} else {
		return_value = nil
	}
	return nil
}

func (n ArrayNode) evaluate(scope *Scope) interface{} {
	var array []interface{}
	for _, e := range n.items {
		array = append(array, e.evaluate(scope))
	}
	return array
}

func (n FunctionCallNode) evaluate(scope *Scope) interface{} {
	if isIntrinsic(n.name) {
		return handleIntrinsic(n, scope)
	} else {
		function := scope.getFromScope(n.name)

		if _, isFunc := function.(Function); isFunc {
			newScope := createScope(scope)
			for i, a := range function.(Function).params {
				newScope.addToScope(a, n.args[i].evaluate(scope))
			}
			for _, b := range function.(Function).body {
				b.evaluate(&newScope)
				if exit_function {
					exit_function = false
					return return_value
				}
			}
			return nil
		} else if _, isFoldNode := function.(FoldNode); isFoldNode {
			foldInst := make(Fold)
			for i, arg := range n.args {
				foldInst[function.(FoldNode).fields[i]] = arg
			}
			return foldInst
		} else if _, isFold := function.(Fold); isFold {
			return function.(Fold)[n.args[0].String()]
		} else if _, isArray := function.([]interface{}); isArray {
			return function.([]interface{})[n.args[0].evaluate(scope).(int)]
		}
	}
	return nil
}

func (n IdentifierNode) evaluate(scope *Scope) interface{} {
	return scope.getFromScope(n.name)
}

func (n FoldNode) evaluate(scope *Scope) interface{} {
	scope.addToScope(n.name, n)
	return nil
}
