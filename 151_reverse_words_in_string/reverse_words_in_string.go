package reverse_words_in_string

import "strings"

func reverseWords(s string) string {
	ans := strings.Builder{}
	left, right := len(s)-1, len(s)-1

	for left >= 0 {
		if s[left] == ' ' {
			if right != left {
				ans.WriteString(s[left+1 : right+1])
				ans.WriteByte(' ')
			}

			left--
			right = left
		} else if left == 0 {
			ans.WriteString(s[0 : right+1])
			break
		} else {
			left--
		}
	}

	if str := ans.String(); str[len(str)-1] == ' ' {
		return str[:len(str)-1]
	} else {
		return str
	}
}
