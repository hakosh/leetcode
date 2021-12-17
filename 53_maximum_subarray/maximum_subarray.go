package maximum_subarray

import "math"

func maxSubArray(nums []int) int {
	best := math.MinInt32
	current := 0

	for _, n := range nums {
		current = max(current+n, n)
		best = max(best, current)
	}

	return best
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
