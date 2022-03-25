package two_city_scheduling

import (
	"sort"
)

func twoCitySchedCost(costs [][]int) int {
	a := make([][]int, 0, len(costs))
	b := make([][]int, 0, len(costs))

	totalCost := 0

	for _, cost := range costs {
		if cost[0] < cost[1] {
			totalCost += cost[0]
			a = append(a, []int{cost[0], cost[1], cost[1] - cost[0]})
		} else {
			totalCost += cost[1]
			b = append(b, []int{cost[0], cost[1], cost[0] - cost[1]})
		}
	}

	longer, shorter := a, b
	if len(a) < len(b) {
		longer, shorter = b, a
	}

	sort.Slice(longer, func(i, j int) bool {
		return longer[i][2] > longer[j][2]
	})

	for i := 1; i <= (len(costs)/2)-len(shorter); i++ {
		top := longer[len(longer)-i]
		totalCost += abs(top[0] - top[1])
	}

	return totalCost
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
