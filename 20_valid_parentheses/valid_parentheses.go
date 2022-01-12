package valid_parentheses

func isValid(s string) bool {
	stack := make([]int8, 0)

	for _, p := range s {
		p := int8(p)

		if p == '(' || p == '{' || p == '[' {
			stack = append(stack, p)
		} else if len(stack) == 0 {
			return false
		} else {
			l := stack[len(stack)-1]

			if (p == ')' && l == '(') || (p == ']' && l == '[') || (p == '}' && l == '{') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
