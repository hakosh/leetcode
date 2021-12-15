package intersection_of_two_arrays_ii

import (
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	intr := make([]int, 0, 3)

	sort.Ints(nums1)
	sort.Ints(nums2)

	ptr1 := 0
	ptr2 := 0

	for {
		if ptr1 >= len(nums1) || ptr2 >= len(nums2) {
			break
		} else if nums1[ptr1] == nums2[ptr2] {
			intr = append(intr, nums1[ptr1])
			ptr1++
			ptr2++
		} else if nums1[ptr1] < nums2[ptr2] {
			ptr1++
		} else {
			ptr2++
		}
	}

	return intr
}
