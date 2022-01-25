package min_cost_of_climbing_stairs

var m map[int]int

func minCostClimbingStairs(cost []int) int {
	downOne := cost[0]
	downTwo := cost[1]

	for i := 2; i < len(cost); i++ {
		downOne, downTwo = downTwo, min(downTwo, downOne)+cost[i]
	}

	return min(downTwo, downOne)
}

func minCostClimbingStairsR(cost []int) int {
	m = make(map[int]int, len(cost))

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
