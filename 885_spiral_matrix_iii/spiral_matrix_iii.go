package spiral_matrix_iii

var directions = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	cells := make([][]int, 0, rows*cols)

	dir := directions[0]
	r, c := rStart, cStart

	curr := 1
	turns := 0
	steps := 2

	for len(cells) < cap(cells) {
		if r < rows && r >= 0 && c < cols && c >= 0 {
			cells = append(cells, []int{r, c})
		}

		r += dir[0]
		c += dir[1]
		curr++

		if curr == steps {
			turns++
			curr = 1
			dir = directions[turns%4]

			if turns%2 == 0 {
				steps++
			}
		}
	}

	return cells
}
