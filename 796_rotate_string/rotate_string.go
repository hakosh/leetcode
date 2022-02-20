package rotate_string

import "strings"

func rotateString(s string, goal string) bool {
	return strings.Index(s+s, goal) != -1
}
