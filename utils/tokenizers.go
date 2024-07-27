package utils

import (
	"strings"
	"unicode"
)

func TokenizeEquation(equation string) []string {
	var tokens []string
	var number strings.Builder
	insideNumber := false

	for i, ch := range equation {
		if unicode.IsDigit(ch) || ch == '.' {
			if !insideNumber {
				insideNumber = true
			}
			number.WriteRune(ch)
		} else {
			if insideNumber {
				tokens = append(tokens, number.String())
				number.Reset()
				insideNumber = false
			}

			if ch == '-' {
				if i == 0 || !unicode.IsDigit(rune(equation[i-1])) && equation[i-1] != ')' {
					number.WriteRune(ch)
					insideNumber = true
				} else {
					tokens = append(tokens, string(ch))
				}
			} else if ch != ' ' {
				if insideNumber {
					tokens = append(tokens, number.String())
					number.Reset()
					insideNumber = false
				}
				tokens = append(tokens, string(ch))
			}
		}
	}

	if insideNumber {
		tokens = append(tokens, number.String())
	}

	return tokens
}
