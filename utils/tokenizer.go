package utils

import (
	"fmt"
	"strings"
	"unicode"
)

func TokenizeEquation(e string) ([]string, error) {
	var tokens []string
	var number strings.Builder
	e = strings.ReplaceAll(e, " ", "")

	isNegativeSign := func(i int, prevChar rune) bool {
		return i == 0 || prevChar == '(' || (!unicode.IsNumber(prevChar) && prevChar != ')')
	}

	for i, chr := range e {
		switch {
		case unicode.IsNumber(chr) || chr == '.':
			number.WriteRune(chr)
		case chr == '-' && (i == 0 || isNegativeSign(i, rune(e[i-1]))):
			number.WriteRune(chr)
		case chr == '(' || chr == ')' || chr == '+' || chr == '-' || chr == '*' || chr == '/' || chr == '^':
			if number.Len() > 0 {
				tokens = append(tokens, number.String())
				number.Reset()
			}
			tokens = append(tokens, string(chr))
		default:
			return nil, fmt.Errorf("unexpected character: %c", chr)
		}
	}

	if number.Len() > 0 {
		tokens = append(tokens, number.String())
	}

	return tokens, nil
}
