package number_of_islands

type Point struct {
	row int
	col int
}

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	islands := 0

	for i, row := range grid {
		for j, col := range row {
			if col == '0' {
				continue
			}

			islands++
			queue := []Point{{i, j}}

			for len(queue) > 0 {
				var tmp []Point

				for _, point := range queue {
					neighbors := []Point{
						{max(0, point.row-1), point.col},
						{min(m-1, point.row+1), point.col},
						{point.row, max(0, point.col-1)},
						{point.row, min(n-1, point.col+1)},
					}

					for _, point := range neighbors {
						if grid[point.row][point.col] == '1' {
							grid[point.row][point.col] = '0'
							tmp = append(tmp, point)
						}
					}
				}

				queue = tmp
			}
		}
	}

	return islands
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
