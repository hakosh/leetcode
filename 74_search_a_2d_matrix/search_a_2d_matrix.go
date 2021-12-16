package search_a_2d_matrix

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	m := len(matrix[0])

	from, to := 0, n*m-1

	for from <= to {
		mid := (from + to) / 2
		check := matrix[mid/m][mid%m]

		if check == target {
			return true
		} else if check > target {
			to = mid - 1
		} else {
			from = mid + 1
		}
	}

	return false
}
