package utils

import (
	"strings"
	"unicode"
)

func TokenizeEquation(equation string) []string {
	var tokens []string
	var number strings.Builder

	for _, ch := range equation {
		if unicode.IsDigit(ch) || ch == '.' {
			number.WriteRune(ch)
		} else {
			if number.Len() > 0 {
				tokens = append(tokens, number.String())
				number.Reset()
			}
			if ch != ' ' {
				tokens = append(tokens, string(ch))
			}
		}
	}
	if number.Len() > 0 {
		tokens = append(tokens, number.String())
	}

	return tokens
}
