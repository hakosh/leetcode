package __reverse_integer

import (
	"math"
)

func reverse(x int) int {
	isNegative := x < 0
	x = int(math.Abs(float64(x)))
	var reversed int

	for x != 0 {
		digit := x % 10
		x = x / 10

		if reversed > math.MaxInt32/10 {
			return 0
		}

		reversed *= 10

		if reversed > math.MaxInt32-digit {
			return 0
		}

		reversed += digit
	}

	if isNegative {
		return -reversed
	}

	return reversed
}
