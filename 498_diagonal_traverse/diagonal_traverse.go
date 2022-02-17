package diagonal_traverse

func findDiagonalOrder(mat [][]int) []int {
	m := len(mat)
	n := len(mat[0])

	nums := make([]int, m*n)

	dirs := [2][2]int{{-1, 1}, {1, -1}}
	dir := dirs[0]
	row, col := 0, 0

	for i := 0; i < m*n; i++ {
		nums[i] = mat[row][col]
		row, col = row+dir[0], col+dir[1]

		if col > n-1 { // right
			row, col = row+2, col-1
			dir = dirs[1]
		} else if row < 0 { // top
			row = row + 1
			dir = dirs[1]
		} else if row > m-1 { // bottom
			row, col = row-1, col+2
			dir = dirs[0]
		} else if col < 0 { // left
			col = col + 1
			dir = dirs[0]
		}
	}

	return nums
}
