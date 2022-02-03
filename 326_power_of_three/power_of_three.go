package power_of_three

import (
	"math"
)

const epsilon = 0.00000000001

func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}

	x := math.Log10(float64(n)) / math.Log10(3)

	return x-float64(int(x)) < epsilon
}
