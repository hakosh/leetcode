package __palindrome_number

import (
	"strconv"
)

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	chars := strconv.Itoa(x)
	maxRounds := len(chars)/2 + 1

	for i := 0; i < maxRounds; i += 1 {
		if chars[i] != chars[len(chars)-i-1] {
			return false
		}
	}

	return true
}

func IsPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	var reversed int
	i := x

	for i != 0 {
		reversed = reversed*10 + i%10
		i /= 10
	}

	return x == reversed
}
