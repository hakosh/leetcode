package spiral_matrix

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])

	nums := make([]int, 0, m*n)

	offset := 0
	row, col := 0, -1

	for len(nums) < m*n {
		// left to right
		for i := col + 1; i < n-offset; i++ {
			col = i
			nums = append(nums, matrix[row][i])
		}

		// down
		for i := row + 1; i < m-offset; i++ {
			row = i
			nums = append(nums, matrix[i][col])
		}

		if m-2*offset == 1 {
			break
		}

		// right to left
		for i := col - 1; i >= offset; i-- {
			col = i
			nums = append(nums, matrix[row][i])
		}

		if n-2*offset == 1 {
			break
		}

		// up
		for i := row - 1; i > offset; i-- {
			row = i
			nums = append(nums, matrix[i][col])
		}

		offset++
	}

	return nums
}
