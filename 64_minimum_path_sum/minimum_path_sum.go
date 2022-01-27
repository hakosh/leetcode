package minimum_path_sum

var cache [][]int

func minPathSum(grid [][]int) int {
	cache = make([][]int, len(grid))
	for i := range cache {
		cache[i] = make([]int, len(grid[0]))
	}

	return dp(grid, 0, 0)
}

func dp(grid [][]int, i, j int) int {
	m := len(grid) - 1
	n := len(grid[0]) - 1

	if i == m && j == n {
		return grid[i][j]
	}

	if cache[i][j] == 0 {
		minRoute := 0

		if i < m && j < n {
			minRoute = min(dp(grid, i+1, j), dp(grid, i, j+1))
		} else if i < m {
			minRoute = dp(grid, i+1, j)
		} else {
			minRoute = dp(grid, i, j+1)
		}

		cache[i][j] = grid[i][j] + minRoute
	}

	return cache[i][j]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
