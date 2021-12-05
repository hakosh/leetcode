package find_all_numbers_disappeared_in_an_array

func findDisappearedNumbers(nums []int) []int {
	miss := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		v := abs(nums[i]) - 1
		nums[v] = -abs(nums[v])
	}

	for i, v := range nums {
		if v > 0 {
			miss = append(miss, i+1)
		}
	}

	return miss
}

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}
