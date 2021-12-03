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
		a := nums[left] * nums[left]
		b := nums[right] * nums[right]

		if a > b {
			squares[i] = a
			left++
		} else {
			squares[i] = b
			right--
		}
	}

	return squares
}
