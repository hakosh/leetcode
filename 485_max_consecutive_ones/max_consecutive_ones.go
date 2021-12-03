package max_consecutive_ones

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	current := 0

	for _, v := range nums {
		if v == 0 {
			current = 0
		} else {
			current++

			if current > max {
				max = current
			}
		}
	}

	return max
}
