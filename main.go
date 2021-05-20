package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	script, err := ioutil.ReadFile("./test.dabu")
	if err != nil {
		panic(err)
	}

	// fmt.Print("======== Lexer ========\n")

	// for _, t := range lex(script) {
	// 	fmt.Print(t.String())
	// }

	fmt.Print("======== Parser ========\n")

	_, nodes := parse(lex(script))
	for _, n := range nodes {
		fmt.Print(n)
	}
}
