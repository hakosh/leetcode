package find_first_and_last_position_of_element_in_sorted_array

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 || target > nums[len(nums)-1] || target < nums[0] {
		return []int{-1, -1}
	}

	idx := search(nums, target, 0, len(nums))
	left, right := idx, idx

	for left > 0 {
		if p := search(nums, target, 0, left-1); p != -1 {
			left = p
		} else {
			break
		}
	}

	for right > 0 && right < len(nums)-1 {
		if p := search(nums, target, right+1, len(nums)); p != -1 {
			right = p
		} else {
			break
		}
	}

	return []int{left, right}
}

func search(nums []int, target, left, right int) int {
	mid := (left + right) / 2

	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return search(nums, target, left, mid)
	} else {
		return search(nums, target, mid+1, right)
	}
}
