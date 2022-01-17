package pow_x_n

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	} else if n == 0 {
		return 1
	}

	r := myPow(x, n/2)

	if n%2 == 1 {
		r = r * r * x
	} else {
		r = r * r
	}

	return r
}
