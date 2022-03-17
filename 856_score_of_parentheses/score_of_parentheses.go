package score_of_parentheses

import "math"

func scoreOfParentheses(s string) int {
	score := 0
	depth := 0

	for i, p := range s {
		if p == '(' {
			depth++
		} else if s[i-1] == '(' {
			depth--
			score += int(math.Pow(2, float64(depth)))
		} else {
			depth--
		}
	}

	return score
}
