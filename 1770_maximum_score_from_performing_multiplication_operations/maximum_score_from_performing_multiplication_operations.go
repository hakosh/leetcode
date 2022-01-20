package maximum_score_from_performing_multiplication_operations

var cache [][]int

func maximumScore(nums, multipliers []int) int {
	dp := make([][]int, len(multipliers)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(multipliers)+1)
	}

	for i := len(multipliers) - 1; i >= 0; i-- {
		for left := i; left >= 0; left-- {
			m := multipliers[i]
			right := len(nums) - 1 - (i - left)

			dp[i][left] = max(
				nums[left]*m+dp[i+1][left+1],
				nums[right]*m+dp[i+1][left],
			)
		}
	}

	return dp[0][0]
}

func maximumScoreR(nums []int, multipliers []int) int {
	cache = make([][]int, len(multipliers))
	for i, _ := range cache {
		cache[i] = make([]int, len(multipliers))
	}

	return dp(nums, 0, multipliers)
}

func dp(nums []int, left int, multipliers []int) int {
	m := multipliers[0]

	if len(multipliers) == 0 {
		return 0
	}

	if len(multipliers) == 1 {
		return max(nums[0]*m, nums[len(nums)-1]*m)
	}

	in := len(nums) - 1
	im := len(multipliers) - 1

	if cache[im][left] == 0 {
		cache[im][left] = max(
			dp(nums[1:], left+1, multipliers[1:])+nums[0]*m,
			dp(nums[:in], left, multipliers[1:])+nums[in]*m,
		)
	}

	return cache[im][left]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
