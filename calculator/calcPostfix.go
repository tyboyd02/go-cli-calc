package calculator

import (
	"fmt"
	"strconv"

	"github.com/tyboyd02/go-cli-calc/utils"
)

func SolvePostfix(e []string) (float64, error) {
	stack := []string{}

	for _, value := range e {
		switch {
		case utils.IsNumeric(value):
			stack = append(stack, value)
		default:
			if len(stack) < 2 {
				return 0, fmt.Errorf("not enough numbers to compleate expression, Operator: '%s', Stack: %v", value, stack)
			}

			lastIndex := len(stack) - 1

			num1, err := strconv.ParseFloat(stack[lastIndex], 64)
			if err != nil {
				return 0, fmt.Errorf("unable to convert %s to float64", stack[lastIndex])
			}
			num2, err := strconv.ParseFloat(stack[lastIndex-1], 64)
			if err != nil {
				return 0, fmt.Errorf("unable to convert %s to float64", stack[lastIndex])
			}

			stack = (stack)[:lastIndex-1]

			sum, err := calculate(num2, num1, value)
			if err != nil {
				return 0, err
			}
			stack = append(stack, strconv.FormatFloat(sum, 'E', -1, 64))
		}
	}

	if len(stack) > 1 {
		return 0, fmt.Errorf("stack is greater than 1 with no operators left, stack: %v", stack)
	}

	total, err := strconv.ParseFloat(stack[0], 64)
	if err != nil {
		return 0, fmt.Errorf("unable to convert '%s' to float64", stack[0])
	}
	return total, nil
}
