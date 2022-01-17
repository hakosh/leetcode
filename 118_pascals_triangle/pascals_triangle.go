package pascals_triangle

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}

	tri := generate(numRows - 1)
	row := make([]int, numRows)

	prev := tri[len(tri)-1]

	row[0] = 1
	for i := 1; i < numRows-1; i++ {
		row[i] = prev[i-1] + prev[i]
	}
	row[numRows-1] = 1

	return append(tri, row)
}
