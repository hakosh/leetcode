package paint_house_iii

import (
	"math"
)

var cache [][][]int

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	cache = make([][][]int, m)
	for i := range cache {
		cache[i] = make([][]int, n+1)

		for j := range cache[i] {
			cache[i][j] = make([]int, target+2)
		}
	}

	if res := dp(houses, cost, target+1, 0, n); res == math.MaxInt32 {
		return -1
	} else {
		return res
	}
}

func dp(houses []int, cost [][]int, nhoods int, house int, lastColor int) int {
	if nhoods == 0 {
		return math.MaxInt32
	}

	if house >= len(houses) {
		if nhoods > 1 {
			return math.MaxInt32
		} else {
			return 0
		}
	}

	if cache[house][lastColor][nhoods] == 0 {
		best := math.MaxInt32

		if houses[house] == 0 {
			for i, price := range cost[house] {
				hoodsLeft := nhoods
				if i != lastColor {
					hoodsLeft--
				}

				sum := price + dp(houses, cost, hoodsLeft, house+1, i)
				best = min(sum, best)
			}
		} else {
			paint := houses[house] - 1

			hoodsLeft := nhoods
			if paint != lastColor {
				hoodsLeft--
			}

			best = dp(houses, cost, hoodsLeft, house+1, paint)
		}

		cache[house][lastColor][nhoods] = best
	}

	return cache[house][lastColor][nhoods]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
