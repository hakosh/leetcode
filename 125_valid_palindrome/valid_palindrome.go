package valid_palindrome

func isPalindrome(s string) bool {
	head := 0
	tail := len(s) - 1

	for head <= tail {
		if isInvalid(s[head]) {
			head++
			continue
		}
		if isInvalid(s[tail]) {
			tail--
			continue
		}

		if isMatch(s[head], s[tail]) {
			head++
			tail--
		} else {
			return false
		}
	}

	return true
}

func isInvalid(c uint8) bool {
	return (c < 'A' || c > 'Z') && (c < 'a' || c > 'z') && (c < '0' || c > '9')
}

func isMatch(a, b uint8) bool {
	if a >= 'A' && a <= 'Z' {
		a += 32
	}

	if b >= 'A' && b <= 'Z' {
		b += 32
	}

	return a == b
}
