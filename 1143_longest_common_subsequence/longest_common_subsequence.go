package longest_common_subsequence

var cache [][]int

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text2)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(text1)+1)
	}

	for i := len(text2) - 1; i >= 0; i-- {
		for j := len(text1) - 1; j >= 0; j-- {
			if text2[i] == text1[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[0][0]
}

func longestCommonSubsequenceR(text1 string, text2 string) int {
	cache = make([][]int, len(text2)+1)
	for i, _ := range cache {
		cache[i] = make([]int, len(text1)+1)
	}

	return dp(text1, text2)
}

func dp(text1 string, text2 string) int {
	i := len(text2)
	j := len(text1)

	if i == 0 || j == 0 {
		return 0
	}

	if cache[i][j] == 0 {
		if text1[0] == text2[0] {
			cache[i][j] = dp(text1[1:], text2[1:]) + 1
		} else {
			cache[i][j] = max(
				dp(text1[1:], text2),
				dp(text1, text2[1:]),
			)
		}
	}

	return cache[i][j]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
