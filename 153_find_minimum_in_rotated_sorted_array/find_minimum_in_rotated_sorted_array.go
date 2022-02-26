package find_minimum_in_rotated_sorted_array

func findMin(nums []int) int {
	pivot := findPivot(nums)
	return nums[pivot]
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
