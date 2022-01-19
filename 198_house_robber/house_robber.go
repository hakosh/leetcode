package house_robber

var mem map[int]int

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums)+1)
	dp[len(nums)-1] = nums[len(nums)-1]
	dp[len(nums)-2] = nums[len(nums)-2]

	for i := len(nums) - 2; i >= 0; i-- {
		dp[i] = max(nums[i]+dp[i+2], dp[i+1])
	}

	return dp[0]
}

func robRec(nums []int) int {
	mem = make(map[int]int, len(nums))
	return dp(nums)
}

func dp(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	}

	li := len(nums) - 1

	if v, ok := mem[li]; ok {
		return v
	} else {
		r := max(nums[li]+dp(nums[:li-1]), dp(nums[:li]))
		mem[li] = r
		return r
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
