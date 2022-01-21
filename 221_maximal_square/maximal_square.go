package maximal_square

func maximalSquare(matrix [][]byte) int {
	cache := make([][]int, len(matrix)+1)
	for i, _ := range cache {
		cache[i] = make([]int, len(matrix[0])+1)
	}

	var m int

	for i := len(matrix) - 1; i >= 0; i-- {
		for j := len(matrix[0]) - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				cache[i][j] = 1 + min(cache[i][j+1], min(cache[i+1][j], cache[i+1][j+1]))

				if cache[i][j] > m {
					m = cache[i][j]
				}
			}
		}
	}

	return m * m
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
