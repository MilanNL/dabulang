package main

import (
	"fmt"
	"regexp"
)

func lex(script []byte) (tokens []Token) {
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
			tokens = append(tokens, Token{TOKEN_NUMBER, match})
			script = script[len(match):]
		} else if match := identifierRegex.Find(script); match != nil {
			tokens = append(tokens, Token{TOKEN_IDENTIFIER, match})
			script = script[len(match):]
		} else if match := bracketRegex.Find(script); match != nil {
			tokens = append(tokens, Token{TOKEN_BRACKET, match})
			script = script[1:]
		} else if match := stringLiteralRegex.Find(script); match != nil {
			tokens = append(tokens, Token{TOKEN_STRINGLITERAL, match})
			script = script[len(match):]
		} else if match := separatorRegex.Find(script); match != nil {
			tokens = append(tokens, Token{TOKEN_SEPARATOR, match})
			script = script[len(match):]
		} else if match := operatorRegex.Find(script); match != nil {
			tokens = append(tokens, Token{TOKEN_OPERATOR, match})
			script = script[len(match):]
		} else {
			panic("Unknown token")
		}
	}
	return
}

func iswhitespace(c byte) bool {
	return c == ' ' || c == '\r' || c == '\n'
}

func (t Token) String() (s string) {
	switch t.tokenType {
	case TOKEN_BRACKET:
		s = fmt.Sprintf("Bracket: %s\n", t.value)
	case TOKEN_IDENTIFIER:
		s = fmt.Sprintf("Identifier: %s\n", t.value)
	case TOKEN_OPERATOR:
		s = fmt.Sprintf("Operator: %s\n", t.value)
	case TOKEN_STRINGLITERAL:
		s = fmt.Sprintf("String literal: %s\n", t.value)
	case TOKEN_NUMBER:
		s = fmt.Sprintf("Number: %s\n", t.value)
	case TOKEN_SEPARATOR:
		s = fmt.Sprintf("Separator: %s\n", t.value)
	}
	return
}

const (
	TOKEN_IDENTIFIER    = iota
	TOKEN_NUMBER        = iota
	TOKEN_BRACKET       = iota
	TOKEN_STRINGLITERAL = iota
	TOKEN_SEPARATOR     = iota
	TOKEN_OPERATOR      = iota
)

type Token struct {
	tokenType int
	value     []byte
}
