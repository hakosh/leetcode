package find_minimum_in_rotated_sorted_array_ii

func findMin(nums []int) int {
	pivot := findPivot(nums)
	return nums[pivot]
}

func findPivot(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[left] {
			right = mid
		} else {
			right--
		}
	}

	return left
}
