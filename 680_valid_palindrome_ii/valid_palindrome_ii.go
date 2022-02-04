package valid_palindrome_ii

func validPalindrome(s string) bool {
	return isValid(s, 1)
}

func isValid(s string, skip int) bool {
	head, tail := 0, len(s)-1

	for head < tail {
		if s[head] == s[tail] {
			head++
			tail--
		} else if skip > 0 {
			return isValid(s[head+1:tail+1], skip-1) || isValid(s[head:tail], skip-1)
		} else {
			return false
		}
	}

	return true
}
