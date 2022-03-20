package campus_bikes

import (
	"sort"
)

func assignBikes(workers [][]int, bikes [][]int) []int {
	bikeSet := make([]bool, len(bikes))
	workerSet := make([]bool, len(workers))

	distances := make([][3]int, 0, len(bikes)*len(workers))

	for i, worker := range workers {
		for j, bike := range bikes {
			entry := [3]int{i, j, manhattan(worker, bike)}
			distances = append(distances, entry)
		}
	}

	sort.SliceStable(distances, func(i, j int) bool {
		if distances[i][2] != distances[j][2] {
			return distances[i][2] < distances[j][2]
		}

		if distances[i][0] != distances[j][0] {
			return distances[i][0] < distances[i][0]
		}

		return distances[i][1] < distances[j][1]
	})

	ans := make([]int, len(workers))

	for _, distance := range distances {
		w, b := distance[0], distance[1]

		if !workerSet[w] && !bikeSet[b] {
			ans[w] = b
			bikeSet[b] = true
			workerSet[w] = true
		}
	}

	return ans
}

func manhattan(a []int, b []int) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
