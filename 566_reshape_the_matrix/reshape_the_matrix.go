package reshape_the_matrix

func matrixReshape(mat [][]int, r int, c int) [][]int {
	n := len(mat)
	m := len(mat[0])

	if m*n != r*c {
		return mat
	}

	out := make([][]int, r)

	for i := 0; i < r; i++ {
		out[i] = make([]int, c)
	}

	for i := 0; i < m*n; i++ {
		i1, j1 := i/m, i%m
		i2, j2 := i/c, i%c

		out[i2][j2] = mat[i1][j1]
	}

	return out
}
