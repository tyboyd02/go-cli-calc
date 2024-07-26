package calculator

import (
	"fmt"
	"math"
)

func calculate(num1, num2 float64, operator string) (float64, error) {
	var result float64

	switch {
	case operator == "^":
		result = math.Pow(num1, num2)
	case operator == "*":
		result = num1 * num2
	case operator == "/":
		if num2 == 0 {
			return 0, fmt.Errorf("can't divide by 0, tried %.2f/%.2f", num1, num2)
		}
		result = num1 / num2
	case operator == "+":
		result = num1 + num2
	case operator == "-":
		result = num1 - num2
	default:
		return 0, fmt.Errorf("'%s' is not an allowed operator", operator)
	}

	return result, nil
}
