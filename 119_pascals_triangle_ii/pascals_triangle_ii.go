package pascals_triangle_ii

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	prv := getRow(rowIndex - 1)
	row := make([]int, rowIndex+1)

	row[0] = 1
	for i := 1; i < rowIndex; i++ {
		row[i] = prv[i-1] + prv[i]
	}
	row[rowIndex] = 1

	return row
}
