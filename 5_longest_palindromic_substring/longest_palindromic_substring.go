package longest_palindromic_substring

func longestPalindrome(s string) string {
	longest := string(s[0])

	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s)+1)
	}

	for to := 0; to <= len(s); to++ {
		for from := 0; from < to; from++ {
			if to-from > 2 {
				dp[from][to] = s[from] == s[to-1] && dp[from+1][to-1]
			} else {
				dp[from][to] = s[from] == s[to-1]
			}

			if dp[from][to] && to-from > len(longest) {
				longest = s[from:to]
			}
		}
	}

	return longest
}
