package smallest_string_with_a_given_numeric_value

func getSmallestString(n int, k int) string {
	word := make([]byte, n)

	for pos := n - 1; pos >= 0; pos-- {
		add := min(k-pos, 26)
		word[pos] = byte('a' + add - 1)
		k -= add
	}

	return string(word)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
