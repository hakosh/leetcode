package word_break

// BOTTOM UP

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 0; i < len(s); i++ {
		for _, word := range wordDict {
			at := i + len(word)

			if at > len(dp) || len(word) > len(s[i:]) {
				continue
			}

			if !dp[at] {
				dp[at] = s[i:i+len(word)] == word && dp[i]
			}
		}
	}

	return dp[len(s)]
}

// TOP DOWN

var cache []int

func wordBreakR(s string, wordDict []string) bool {
	cache = make([]int, len(s))
	return dp(s, wordDict, 0)
}

func dp(s string, dict []string, from int) bool {
	if len(s[from:]) == 0 {
		return true
	}

	if cache[from] != 0 {
		return cache[from] == 2
	}

	result := false

	for _, word := range dict {
		if len(word) > len(s[from:]) {
			continue
		}

		result = s[from:from+len(word)] == word && dp(s, dict, from+len(word))

		if result {
			break
		}
	}

	if result {
		cache[from] = 2
	} else {
		cache[from] = 1
	}

	return result
}
