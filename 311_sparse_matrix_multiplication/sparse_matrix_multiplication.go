package sparse_matrix_multiplication

func multiply(mat1 [][]int, mat2 [][]int) [][]int {
	n := len(mat1)
	m := len(mat2[0])

	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			mat[i][j] = mul(mat1, i, mat2, j)
		}
	}

	return mat
}

func mul(mat1 [][]int, row int, mat2 [][]int, col int) int {
	res := 0

	for i := 0; i < len(mat1[row]); i++ {
		res += mat1[row][i] * mat2[i][col]
	}

	return res
}
