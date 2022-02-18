package merge_intervals

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)

	open := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= open[1] {
			open[1] = max(open[1], intervals[i][1])
		} else {
			merged = append(merged, open)
			open = intervals[i]
		}
	}

	return append(merged, open)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
