package min_cost_of_climbing_stairs

import "fmt"

var m map[int]int

func minCostClimbingStairs(cost []int) int {
	downTwo, downOne := 0, 0

	for i := 2; i < len(cost)+1; i++ {
		downOne, downTwo = min(downTwo+cost[i-2], downOne+cost[i-1]), downOne
	}

	return downOne
}

func minCostClimbingStairsR(cost []int) int {
	m = make(map[int]int, len(cost))

	defer func() {
		fmt.Println(m)
	}()

	m[0] = cost[0]
	m[1] = cost[1]

	return min(dp(cost), dp(cost[:len(cost)-1]))
}

func dp(cost []int) int {
	li := len(cost) - 1

	if v, ok := m[li]; ok {
		return v
	} else {
		r := min(dp(cost[:li]), dp(cost[:li-1])) + cost[li]
		m[li] = r

		return r
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
