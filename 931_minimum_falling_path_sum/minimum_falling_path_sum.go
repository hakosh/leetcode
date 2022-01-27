package minimum_falling_path_sum

import "math"

var cache [][]int

func minFallingPathSum(matrix [][]int) int {
	cache = make([][]int, len(matrix))
	for i := range cache {
		cache[i] = make([]int, len(matrix[0]))
	}

	sum := math.MaxInt32

	for j := 0; j < len(matrix[0]); j++ {
		if minSum := dp(matrix, 0, j); minSum < sum {
			sum = minSum
		}
	}

	return sum
}

func dp(matrix [][]int, i, j int) int {
	m := len(matrix) - 1
	n := len(matrix[0]) - 1

	if j < 0 || j > n {
		return math.MaxInt32
	}

	if m == i {
		return matrix[i][j]
	}

	if cache[i][j] == 0 {
		cache[i][j] = matrix[i][j] + min(dp(matrix, i+1, j), min(dp(matrix, i+1, j-1), dp(matrix, i+1, j+1)))
	}

	return cache[i][j]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
