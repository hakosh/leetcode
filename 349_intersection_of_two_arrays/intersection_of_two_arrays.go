package intersection_of_two_arrays

import (
	"math"
	"sort"
)

func intersection(nums1 []int, nums2 []int) []int {
	intr := make([]int, 0, 3)

	sort.Ints(nums1)
	sort.Ints(nums2)

	ptr1 := 0
	ptr2 := 0

	last := math.MinInt32

	for {
		if ptr1 >= len(nums1) || ptr2 >= len(nums2) {
			break
		} else if nums1[ptr1] == nums2[ptr2] {
			if nums1[ptr1] != last {
				intr = append(intr, nums1[ptr1])
				last = nums1[ptr1]
			}

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

func intersection2(nums1 []int, nums2 []int) []int {
	items1 := make(map[int]bool)
	for _, num := range nums1 {
		items1[num] = true
	}

	answer := make([]int, 0)
	items2 := make(map[int]bool)

	for _, num := range nums2 {
		if _, ok := items1[num]; ok {
			if _, ok := items2[num]; ok {
				continue
			}

			answer = append(answer, num)
			items2[num] = true
		}
	}

	return answer
}