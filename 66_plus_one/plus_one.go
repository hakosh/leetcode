package plus_one

func plusOne(digits []int) []int {
	result := make([]int, len(digits)+1)
	copy(result[1:], digits)

	last := len(result) - 1
	overflow := result[last] == 9
	if !overflow {
		result[last]++
	}

	for i := len(result) - 1; i >= 0 && overflow; i-- {
		this := result[i] + 1
		overflow = this > 9

		if overflow {
			result[i] = 0
		} else {
			result[i] = this
		}
	}

	if result[0] == 0 {
		return result[1:]
	}

	return result
}
