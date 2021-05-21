package main

import "fmt"

var intrinsicNames = [...]string{"print"}
var intrinsicImpls = [...](func(FunctionCallNode, *Scope) interface{}){print}

func isIntrinsic(name string) bool {
	for _, n := range intrinsicNames {
		if n == name {
			return true
		}
	}
	return false
}

func handleIntrinsic(node FunctionCallNode, scope *Scope) interface{} {
	for i, name := range intrinsicNames {
		if name == node.name {
			return intrinsicImpls[i](node, scope)
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
