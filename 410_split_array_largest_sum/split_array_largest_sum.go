package split_array_largest_sum

import (
	"math"
)

var cache [][]int
var sumRight []int

func splitArray(nums []int, m int) int {
	cache = make([][]int, m+1)
	for i := range cache {
		cache[i] = make([]int, len(nums))
	}

	sumRight = make([]int, len(nums))
	sumRight[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		sumRight[i] = nums[i] + sumRight[i+1]
	}

	return dp(nums, m, 0)
}

func dp(weights []int, days int, from int) int {
	if days == 1 {
		cache[days][from] = sumRight[from]
		return sumRight[from]
	}

	if weights[from] == 0 && len(weights[from:]) > days {
		return dp(weights, days, from+1)
	}

	if cache[days][from] == 0 {
		best := math.MaxInt32

		for i := from + 1; i <= len(weights)-days+1; i++ {
			first := sumRight[from] - sumRight[i]
			best = min(best, max(first, dp(weights, days-1, i)))

			if first >= best {
				break
			}
		}

		cache[days][from] = best
	}

	return cache[days][from]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
