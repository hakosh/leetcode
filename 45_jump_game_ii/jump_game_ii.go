package jump_game_ii

func jump(nums []int) int {
	jumps, currentJumpEnd, farthest := 0, 0, 0

	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])

		if i == currentJumpEnd {
			jumps++
			currentJumpEnd = farthest
		}
	}

	return jumps
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func jumpDp(nums []int) int {
	dp := make([]int, len(nums))

	for at, jumps := range nums {
		for i := 1; i <= jumps && at+i < len(nums); i++ {
			if dp[at+i] == 0 {
				dp[at+i] = dp[at] + 1
			}
		}
	}

	return dp[len(nums)-1]
}
