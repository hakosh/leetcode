package maximum_sum_circular_subarray

func maxSubarraySumCircular(nums []int) int {
	N := len(nums)
	ans, cur := nums[0], nums[0]

	for i := 1; i < N; i++ {
		cur = nums[i] + max(cur, 0)
		ans = max(ans, cur)
	}

	rightsums := make([]int, N)
	rightsums[N-1] = nums[N-1]
	for i := N - 2; i >= 0; i-- {
		rightsums[i] = rightsums[i+1] + nums[i]
	}

	maxright := make([]int, N)
	maxright[N-1] = nums[N-1]
	for i := N - 2; i >= 0; i-- {
		maxright[i] = max(maxright[i+1], rightsums[i])
	}

	leftsum := 0
	for i := 0; i < len(nums)-2; i++ {
		leftsum += nums[i]
		ans = max(ans, leftsum+maxright[i+2])
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
