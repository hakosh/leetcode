package unique_paths_ii

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	lastRow := make([]int, n+1)
	thisRow := make([]int, n+1)

	lastRow[1] = 1

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if obstacleGrid[i-1][j-1] == 1 {
				continue
			}

			thisRow[j] = lastRow[j] + thisRow[j-1]
		}

		lastRow, thisRow = thisRow, make([]int, n+1)
	}

	return lastRow[n]
}
