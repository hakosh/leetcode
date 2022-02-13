package palindrome_permutation

func canPermutePalindrome(s string) bool {
	count := make(map[rune]int)

	for _, c := range s {
		count[c]++
	}

	mid := false

	for _, v := range count {
		if v%2 == 1 {
			if !mid {
				mid = true
			} else {
				return false
			}
		}
	}

	return true
}
