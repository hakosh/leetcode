package reverse_words_in_string_ii

func reverseWords(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}

	left, right := 0, 0

	for left < len(s) {
		if right == len(s) || s[right] == ' ' {
			for i := left; i < (left+right)/2; i++ {
				s[i], s[left+right-i-1] = s[left+right-i-1], s[i]
			}

			left = right + 1
		}

		right++
	}
}
