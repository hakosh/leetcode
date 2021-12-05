package max_consecutive_ones_ii

func findMaxConsecutiveOnes(nums []int) int {
	max := 0

	zero := 0
	last := 0
	current := 0

	for _, v := range nums {
		if v == 1 {
			current++
		} else {
			zero, last, current = 1, current, 0
		}

		tmp := zero + last + current

		if zero+last+current > max {
			max = tmp
		}
	}

	return max
}
