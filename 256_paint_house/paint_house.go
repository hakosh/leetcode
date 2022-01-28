package paint_house

import (
	"math"
)

// BOTTOM UP

func minCost(costs [][]int) int {
	thisRow := make([]int, 3)
	lastRow := costs[len(costs)-1]

	for i := len(costs) - 2; i >= 0; i-- {
		for j, cost := range costs[i] {
			switch j {
			case 0:
				thisRow[j] = cost + min(lastRow[1], lastRow[2])
			case 1:
				thisRow[j] = cost + min(lastRow[0], lastRow[2])
			case 2:
				thisRow[j] = cost + min(lastRow[0], lastRow[1])
			}
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

// TOP DOWN

var cache [][]int

func minCostR(costs [][]int) int {
	cache = make([][]int, len(costs))
	for i := range cache {
		cache[i] = make([]int, len(costs[0]))
	}

	sum := math.MaxInt32

	for j := 0; j < len(costs[0]); j++ {
		if minSum := dp(costs, 0, j); minSum < sum {
			sum = minSum
		}
	}

	return sum
}

func dp(costs [][]int, i, j int) int {
	m := len(costs) - 1

	if m == i {
		return costs[i][j]
	}

	if cache[i][j] == 0 {
		switch j {
		case 0:
			cache[i][j] = costs[i][j] + min(dp(costs, i+1, 1), dp(costs, i+1, 2))
		case 1:
			cache[i][j] = costs[i][j] + min(dp(costs, i+1, 0), dp(costs, i+1, 2))
		case 2:
			cache[i][j] = costs[i][j] + min(dp(costs, i+1, 0), dp(costs, i+1, 1))
		}
	}

	return cache[i][j]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
