package largest_number_at_least_twice_of_others

func dominantIndex(nums []int) int {
	count := make([]int, 101)
	max := 0

	for i, num := range nums {
		count[num]++
		if num > nums[max] {
			max = i
		}
	}

	if nums[max] == 1 {
		return max
	}

	for i := nums[max] - 1; i > nums[max]/2; i-- {
		if count[i] != 0 {
			return -1
		}
	}

	return max
}
