package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	ans := 0
	dict := make(map[uint8]int)

	for i, j := 0, 0; j < len(s); j++ {
		if v, ok := dict[s[j]]; ok {
			i = max(i, v)
		}

		ans = max(ans, j-i+1)
		dict[s[j]] = j + 1
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
