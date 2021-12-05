package height_checker

import (
	"sort"
)

func heightChecker(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)

	mismatches := 0

	for i, v := range heights {
		if v != expected[i] {
			mismatches++
		}
	}

	return mismatches
}
