package main

import (
	"io/ioutil"
)

func main() {
	script, err := ioutil.ReadFile("./test.dabu")
	if err != nil {
		panic(err)
	}

	_, nodes := parse(lex(script))
	for _, n := range nodes {
		n.evaluate(&global_scope)
	}
}
