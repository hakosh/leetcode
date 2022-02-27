package valid_perfect_square

func isPerfectSquare(num int) bool {
	left, right := 1, num/2

	if num == 1 {
		return true
	}

	for left <= right {
		mid := (left + right) / 2

		r := mid * mid

		if r == num {
			return true
		} else if r < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
