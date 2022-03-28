package search_in_rotated_sorted_array_ii

import "sort"

func search(nums []int, target int) bool {
	for nums[0] == nums[len(nums)-1] {
		if nums[0] == target {
			return true
		}

		if len(nums) <= 2 {
			return false
		}

		nums = nums[1 : len(nums)-1]
	}

	k := findPivot(nums)

	pos := sort.Search(len(nums), func(i int) bool {
		return nums[(i+k)%len(nums)] >= target
	})

	rotatedPos := (pos + k) % len(nums)

	return nums[rotatedPos] == target
}

func findPivot(nums []int) int {
	left, right := 0, len(nums)-1
	pivot := 0

	if nums[0] <= nums[len(nums)-1] {
		return 0
	}

	for left <= right {
		mid := (left + right) / 2

		if mid == left {
			if nums[left] > nums[right] {
				return right
			} else {
				return left
			}
		}

		if nums[mid] > nums[right] {
			left = mid
		} else if nums[mid] < nums[left] {
			right = mid
		}
	}

	return pivot
}
