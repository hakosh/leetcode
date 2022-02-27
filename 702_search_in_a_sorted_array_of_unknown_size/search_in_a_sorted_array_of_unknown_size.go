package search_in_a_sorted_array_of_unknown_size

import (
	"math"
	"sort"
)

type ArrayReader struct {
	items []int
}

func (a *ArrayReader) get(index int) int {
	if index >= len(a.items) {
		return math.MaxInt32
	}

	return a.items[index]
}

func search(reader ArrayReader, target int) int {
	left, right := 0, 1
	for reader.get(right) < target {
		left = right
		right *= 2
	}

	pos := sort.Search(right-left, func(i int) bool {
		return reader.get(left+i) >= target
	})

	if reader.get(left+pos) == target {
		return left + pos
	} else {
		return -1
	}
}
