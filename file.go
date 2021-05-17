package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func check(e error) bool {
	if e != nil {
		panic(e)
	}
	return true
}

func iswhitespace(c byte) bool {
	return c == ' ' || c == '\r' || c == '\n'
}

func main() {
	script, err := ioutil.ReadFile("./test.dabu")
	check(err)

	numberRegex := regexp.MustCompile(`^[0-9]+`)
	identifierRegex := regexp.MustCompile(`^[a-zA-Z][a-zA-Z_0-9]*`)
	bracketRegex := regexp.MustCompile(`^[()\[\]{}]`)
	stringLiteralRegex := regexp.MustCompile(`^".*"`)
	separatorRegex := regexp.MustCompile(`^[.,]`)
	operatorRegex := regexp.MustCompile(`^(==|!=|>=|<=|>|<|=|\+|/|\*|-|\+\+|--|&|&&|\||\|\||<<|>>|!|\^|\^\^|~)`)
	for len(script) > 0 {
		for iswhitespace(script[0]) {
			script = script[1:]
		}

		if match := numberRegex.Find(script); match != nil {
			fmt.Printf("Number: %s\n", match)
			script = script[len(match):]
		} else if match := identifierRegex.Find(script); match != nil {
			fmt.Printf("Identifer: %s\n", match)
			script = script[len(match):]
		} else if match := bracketRegex.Find(script); match != nil {
			fmt.Printf("Bracket: %s\n", match)
			script = script[1:]
		} else if match := stringLiteralRegex.Find(script); match != nil {
			fmt.Printf("String Literal: %s\n", match)
			script = script[len(match):]
		} else if match := separatorRegex.Find(script); match != nil {
			fmt.Printf("Separator: %s\n", match)
			script = script[len(match):]
		} else if match := operatorRegex.Find(script); match != nil {
			fmt.Printf("Operator: %s\n", match)
			script = script[len(match):]
		} else {
			fmt.Print(string(script))
			panic("Unknown token")
		}
	}
}
