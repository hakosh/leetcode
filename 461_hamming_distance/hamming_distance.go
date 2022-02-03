package hamming_distance

func hammingDistance(x int, y int) int {
	n := x ^ y
	distance := 0

	for n != 0 {
		if n%2 == 1 {
			distance++
		}

		n = n >> 1
	}

	return distance
}
