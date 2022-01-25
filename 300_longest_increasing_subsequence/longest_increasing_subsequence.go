package longest_increasing_subsequence

// BOTTOM UP

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i, num := range nums {
		for j := 0; j < i; j++ {
			if num > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	longest := 0

	for _, v := range dp {
		if v > longest {
			longest = v
		}
	}

	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// TOP DOWN

var cache [][]int

func lengthOfLIS2(nums []int) int {
	cache = make([][]int, len(nums))
	for i := range cache {
		cache[i] = make([]int, len(nums))
	}

	return dp(nums, 0, -1)
}

func dp(nums []int, from int, last int) int {
	if from >= len(nums) {
		return 0
	}

	if len(nums[from:]) == 1 {
		if last == -1 || nums[from] > nums[last] {
			return 1
		} else {
			return 0
		}
	}

	if cache[from][last+1] != 0 {
		return cache[from][last+1]
	}

	best := 0

	for i, num := range nums[from:] {
		if last != -1 && num <= nums[from] {
			continue
		}

		if res := 1 + dp(nums, from+i, from+i); res > best {
			best = res
		}
	}

	cache[from][last+1] = best

	return best
}
