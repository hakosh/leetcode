package evaluate_reverse_polish_notation

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0, 32)

	for _, token := range tokens {
		num, err := strconv.Atoi(token)

		if err == nil {
			stack = append(stack, num)
		} else {
			top := len(stack) - 1

			op1 := stack[top-1]
			op2 := stack[top]

			res := 0

			switch token {
			case "+":
				res = op1 + op2
			case "-":
				res = op1 - op2
			case "*":
				res = op1 * op2
			case "/":
				res = op1 / op2
			}

			stack = append(stack[:top-1], res)
		}
	}

	return stack[0]
}
