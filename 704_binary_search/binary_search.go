package binary_search

func search(nums []int, target int) int {
	left, right := 0, len(nums)

	for left != right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return -1
}
