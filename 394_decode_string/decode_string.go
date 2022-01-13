package decode_numbers

import (
	"strings"
)

type Frame struct {
	count int
	text  string
}

func (f *Frame) String() string {
	return strings.Repeat(f.text, f.count)
}

func decodeString(s string) string {
	curr := &Frame{1, ""}
	stack := []*Frame{curr}

	last := '_'

	for _, char := range s {
		if isLetter(char) {
			curr.text += string(char)
		} else if isNumeric(char) {
			val := int(char) - 48

			if isNumeric(last) {
				curr.count = curr.count*10 + val
			} else {
				curr = &Frame{val, ""}
				stack = append(stack, curr)
			}
		} else if char == '[' && !isNumeric(last) {
			curr = &Frame{1, ""}
			stack = append(stack, curr)
		} else if char == ']' {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			curr = stack[len(stack)-1]
			curr.text += top.String()
		}

		last = char
	}

	return stack[0].String()
}

func isLetter(char rune) bool {
	return char >= 'a' && char <= 'z'
}

func isNumeric(char rune) bool {
	return char >= '0' && char <= '9'
}
