package knight_probability_in_chessboard

const moves float64 = 8

var directions = [][2]int{
	{-2, -1},
	{-1, -2},
	{2, -1},
	{1, -2},
	{2, 1},
	{1, 2},
	{-2, 1},
	{-1, 2},
}

/// BOTTOM UP
func makeGrid(n int) [][]float64 {
	grid := make([][]float64, n)
	for i := range grid {
		grid[i] = make([]float64, n)
	}

	return grid
}

func knightProbability(n, k, row, column int) float64 {
	this := makeGrid(n)
	last := makeGrid(n)

	last[row][column] = 1

	for l := 1; l <= k; l++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				all := 0.0

				for _, dir := range directions {
					nRow := i + dir[0]
					nCol := j + dir[1]

					if nRow >= 0 && nRow < n && nCol >= 0 && nCol < n {
						all += last[nRow][nCol]
					}
				}

				this[i][j] = all / moves
			}
		}

		last, this = this, last
	}

	sum := 0.0
	for _, row := range last {
		for _, col := range row {
			sum += col
		}
	}

	return sum
}

/// TOP DOWN

var cache [][][]float64

func knightProbabilityR(n int, k int, row int, column int) float64 {
	cache = make([][][]float64, n)
	for i := range cache {
		cache[i] = make([][]float64, n)

		for j := range cache[i] {
			cache[i][j] = make([]float64, k+1)
		}
	}

	return dp(n, k, row, column)
}

func dp(n, k, row, col int) float64 {
	if k == 0 {
		return 1
	}

	if cache[row][col][k] == 0 {
		all := 0.0

		for _, dir := range directions {
			nRow := row + dir[0]
			nCol := col + dir[1]

			if nRow >= 0 && nRow < n && nCol >= 0 && nCol < n {
				all += dp(n, k-1, nRow, nCol)
			}
		}

		cache[row][col][k] = all / moves
	}

	return cache[row][col][k]
}
