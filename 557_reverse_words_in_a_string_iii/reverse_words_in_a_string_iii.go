package reverse_words_in_a_string_iii

import (
	"strings"
)

func reverseWords(s string) string {
	ans := strings.Builder{}
	left, right := 0, 0

	for right <= len(s) {
		if right == len(s) || s[right] == ' ' {
			writeReverse(&ans, s[left:right])

			if right < len(s) {
				ans.WriteByte(' ')
			}

			left = right + 1
		}

		right++
	}

	return ans.String()
}

func writeReverse(sb *strings.Builder, s string) {
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
}
