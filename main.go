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

	for _, t := range lex(script) {
		fmt.Print(t.String())
	}
}
