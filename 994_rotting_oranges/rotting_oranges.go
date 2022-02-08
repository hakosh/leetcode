package rotting_oranges

var directions = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func orangesRotting(grid [][]int) int {
	fresh := 0
	minutes := 0

	queue := make([][2]int, 0)

	for i, row := range grid {
		for j, col := range row {
			if col == 2 {
				queue = append(queue, [2]int{i, j})
			} else if col == 1 {
				fresh++
			}
		}
	}

	for len(queue) > 0 {
		tmp := make([][2]int, 0)

		for _, orange := range queue {
			for _, dir := range directions {
				row := orange[0] + dir[0]
				col := orange[1] + dir[1]

				if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) || grid[row][col] == 0 {
					continue
				}

				if grid[row][col] == 1 {
					fresh--
					grid[row][col] = 2
					tmp = append(tmp, [2]int{row, col})
				}
			}
		}

		if len(tmp) > 0 {
			minutes++
		}

		queue = tmp
	}

	if fresh == 0 {
		return minutes
	} else {
		return -1
	}
}
