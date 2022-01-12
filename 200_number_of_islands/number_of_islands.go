package number_of_islands

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
			queue := [][2]int{{i, j}}

			for len(queue) > 0 {
				var tmp [][2]int

				for _, point := range queue {
					neighbors := [][2]int{
						{point[0] - 1, point[1]},
						{point[0] + 1, point[1]},
						{point[0], point[1] - 1},
						{point[0], point[1] + 1},
					}

					for _, point := range neighbors {
						if point[0] < 0 || point[0] == m || point[1] < 0 || point[1] == n || grid[point[0]][point[1]] == '0' {
							continue
						}

						grid[point[0]][point[1]] = '0'
						tmp = append(tmp, point)
					}
				}

				queue = tmp
			}
		}
	}

	return islands
}

func numIslandsDFS(grid [][]byte) int {
	islands := 0

	m := len(grid)
	n := len(grid[0])

	for i, row := range grid {
		for j, col := range row {
			if col == '1' {
				islands++
				mark(&grid, i, j, m, n)
			}
		}
	}

	return islands
}

func mark(grid *[][]byte, i, j, m, n int) {
	if i < 0 || j < 0 || i >= m || j >= n || (*grid)[i][j] == '0' {
		return
	}

	(*grid)[i][j] = '0'

	mark(grid, i-1, j, m, n)
	mark(grid, i+1, j, m, n)
	mark(grid, i, j-1, m, n)
	mark(grid, i, j+1, m, n)
}
