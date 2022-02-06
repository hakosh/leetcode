package minimum_remove_to_make_valid_parentheses

import (
	"strings"
)

func minRemoveToMakeValid(s string) string {
	remove := make([]int, 0)
	stack := make([]int, 0)

	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else if c == ')' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				remove = append(remove, i)
			}
		}
	}

	remove = append(remove, stack...)

	if len(remove) == 0 {
		return s
	}

	str := strings.Builder{}
	from := 0
	for _, idx := range remove {
		str.WriteString(s[from:idx])
		from = idx + 1
	}
	str.WriteString(s[from:])

	return str.String()
}
