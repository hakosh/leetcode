package paint_house_ii

import "math"

func minCostII(costs [][]int) int {
	thisRow := make([]int, len(costs[0]))
	lastRow := costs[len(costs)-1]

	for i := len(costs) - 2; i >= 0; i-- {
		for j, cost := range costs[i] {
			thisRow[j] = cost + min(lastRow, j)
		}

		if i > 0 {
			lastRow, thisRow = thisRow, costs[i-1]
		}
	}

	if len(costs) == 1 {
		thisRow = lastRow
	}

	best := math.MaxInt32

	for _, sum := range thisRow {
		if sum < best {
			best = sum
		}
	}

	return best
}

func min(nums []int, skip int) int {
	best := math.MaxInt32

	for i, v := range nums {
		if i == skip {
			continue
		}

		if v < best {
			best = v
		}
	}

	return best
}
