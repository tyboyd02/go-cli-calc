package parser

import (
	"fmt"

	"github.com/tyboyd02/go-cli-calc/utils"
)

func getPresident(s string) int {
	var operators = map[string]int{
		"^": 3,
		"*": 2,
		"/": 2,
		"+": 1,
		"-": 1,
		"(": 0,
	}

	num, ok := operators[s]

	if !ok {
		panic(fmt.Sprintf("Cloud not get president of %s", s))
	}

	return num
}

func transferLastItem(src, dest *[]string) {
	if len(*src) == 0 {
		return
	}
	lastIndex := len(*src) - 1
	item := (*src)[lastIndex]
	*src = (*src)[:lastIndex]
	*dest = append(*dest, item)
}

func ConvertToReversePolishNotation(equation []string) []string {
	operatorStack := []string{}
	outputQueue := []string{}

	for _, token := range equation {
		switch {
		case utils.IsNumeric(token):
			outputQueue = append(outputQueue, token)
		case token == "(":
			operatorStack = append(operatorStack, token)
		case token == ")":
			for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1] != "(" {
				transferLastItem(&operatorStack, &outputQueue)
			}
			operatorStack = operatorStack[:len(operatorStack)-1]
		default:
			for len(operatorStack) > 0 && getPresident(operatorStack[len(operatorStack)-1]) >= getPresident(token) {
				transferLastItem(&operatorStack, &outputQueue)
			}
			operatorStack = append(operatorStack, token)
		}
	}

	for len(operatorStack) > 0 {
		transferLastItem(&operatorStack, &outputQueue)
	}

	return outputQueue
}
