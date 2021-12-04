package sort_array_by_parity

// Optimized for space complexity

func sortArrayByParity(nums []int) []int {
	for i, v := range nums {
		if v%2 == 1 {
			next := nextEven(nums, i)

			if next == -1 {
				break
			}

			nums[i], nums[next] = nums[next], nums[i]
		}
	}

	return nums
}

func nextEven(nums []int, from int) int {
	for i := from; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			return i
		}
	}

	return -1
}
