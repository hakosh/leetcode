package find_k_closest_elements

import (
	"sort"
)

func findClosestElements(arr []int, k int, x int) []int {
	pos := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= x
	})

	left, right := pos-1, pos

	for right-left-1 < k {
		if left == -1 {
			right++
		} else if right == len(arr) {
			left--
		} else {
			if abs(arr[left]-x) <= abs(arr[right]-x) {
				left--
			} else {
				right++
			}
		}
	}

	return arr[left+1 : right]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
