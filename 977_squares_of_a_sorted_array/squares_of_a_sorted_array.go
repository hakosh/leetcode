package squares_of_a_sorted_array

func sortedSquares(nums []int) []int {
	if nums[0] >= 0 {
		for i, v := range nums {
			nums[i] = v * v
		}

		return nums
	}

	squares := make([]int, len(nums))

	left := 0
	right := len(nums) - 1

	for i := right; i >= 0; i-- {
		if abs(nums[left]) > abs(nums[right]) {
			squares[i] = nums[left] * nums[left]
			left++
		} else {
			squares[i] = nums[right] * nums[right]
			right--
		}
	}

	return squares
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
