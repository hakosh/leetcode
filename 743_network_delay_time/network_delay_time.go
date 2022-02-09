package network_delay_time

import (
	"math"
)

func networkDelayTime(times [][]int, n int, k int) int {
	follow := make([][][2]int, n)
	k--

	for _, time := range times {
		follow[time[0]-1] = append(follow[time[0]-1], [2]int{time[1] - 1, time[2]})
	}

	cost := dijkstra(follow, k)

	max := -1
	for _, v := range cost {
		if v == math.MaxInt32 {
			return -1
		} else if v > max {
			max = v
		}
	}

	return max
}

func dijkstra(follow [][][2]int, start int) []int {
	cost := make([]int, len(follow))
	for i := range cost {
		cost[i] = math.MaxInt32
	}
	cost[start] = 0

	next := start
	queue := make(map[int]bool)

	for next != -1 {
		for _, edge := range follow[next] {
			c := cost[next] + edge[1]

			if c < cost[edge[0]] {
				queue[edge[0]] = true
				cost[edge[0]] = c
			}
		}

		next = -1
		for k, _ := range queue {
			if next == -1 || cost[k] < cost[next] {
				next = k
			}
		}
		delete(queue, next)
	}

	return cost
}
