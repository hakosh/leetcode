package rotate_image

func rotate(matrix [][]int) {
	n := len(matrix) - 1

	if n == 0 {
		return
	}

	for j := 0; j < n/2; j++ {
		for i := 0; i < n-j*2; i++ {
			tmp := matrix[j][j+i]

			matrix[j][j+i] = matrix[n-i-j][j]
			matrix[n-i-j][j] = matrix[n-j][n-j-i]
			matrix[n-j][n-j-i] = matrix[j+i][n-j]
			matrix[j+i][n-j] = tmp
		}
	}
}
