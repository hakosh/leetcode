package find_pivot_index

func pivotIndex(nums []int) int {
	left := 0
	right := sum(nums)

	for i, v := range nums {
		right -= v

		if left == right {
			return i
		}

		left += v
	}

	return -1
}

func sum(nums []int) int {
	s := 0
	for _, v := range nums {
		s += v
	}
	return s
}
