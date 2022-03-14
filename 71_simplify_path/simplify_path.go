package simplify_path

import (
	"strings"
)

func simplifyPath(path string) string {
	stack := make([]string, 0)
	parts := strings.Split(path, "/")

	for _, part := range parts {
		if part == "." || len(part) == 0 {
			continue
		} else if part == ".." {
			if len(stack) == 0 {
				continue
			} else {
				stack = stack[0 : len(stack)-1]
			}
		} else {
			stack = append(stack, part)
		}
	}

	return "/" + strings.Join(stack, "/")
}
