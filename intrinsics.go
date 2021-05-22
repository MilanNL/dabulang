package main

import "fmt"

var intrinsicNames = [...]string{"print", "len", "println", "string"}
var intrinsicImpls = [...](func(FunctionCallNode, *Scope) interface{}){print, dabu_len, println, dabu_string}

func isIntrinsic(name string) bool {
	for _, n := range intrinsicNames {
		if n == name {
			return true
		}
	}
	return false
}

func handleIntrinsic(node FunctionCallNode, scope *Scope) interface{} {
	if node.name.getType() == NODE_IDENTIFIER {
		for i, name := range intrinsicNames {
			if name == node.name.String() {
				return intrinsicImpls[i](node, scope)
			}
		}
	}
	return nil
}

func print(node FunctionCallNode, scope *Scope) interface{} {
	for _, arg := range node.args {
		fmt.Print(arg.evaluate(scope))
	}
	return nil
}

func println(node FunctionCallNode, scope *Scope) interface{} {
	for _, arg := range node.args {
		fmt.Println(arg.evaluate(scope))
	}
	return nil
}

func dabu_string(node FunctionCallNode, scope *Scope) interface{} {
	arg := node.args[0].evaluate(scope)
	if _, isNum := arg.(float64); isNum {
		return string(make([]byte, int(arg.(float64))))
	}
	return nil
}

func dabu_len(node FunctionCallNode, scope *Scope) interface{} {
	arg := node.args[0].evaluate(scope)
	if _, isArray := arg.([]interface{}); isArray {
		return len(arg.([]interface{}))
	} else if _, isString := arg.(string); isString {
		return len(arg.(string))
	} else {
		return nil
	}
}
