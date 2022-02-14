package sqrt_x

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	low, high := 2, x/2

	for low <= high {
		mid := (low + high) / 2
		square := mid * mid

		if square == x {
			return mid
		} else if square < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return high
}
