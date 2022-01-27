package unique_paths

// BOTTOM UP

func uniquePaths(m, n int) int {
	lastRow := make([]int, n+1)
	thisRow := make([]int, n+1)

	lastRow[1] = 1

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			thisRow[j] = lastRow[j] + thisRow[j-1]
		}

		lastRow, thisRow = thisRow, make([]int, n+1)
	}

	return lastRow[n]
}

// TOP DOWN

var cache [][]int

func uniquePathsR(m int, n int) int {
	cache = make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
	}

	return dp(m-1, n-1, 0, 0)
}

func dp(m, n, i, j int) int {
	if i > m || j > n {
		return 0
	}

	if i == m && j == n {
		return 1
	}

	if cache[i][j] == 0 {
		cache[i][j] = dp(m, n, i+1, j) + dp(m, n, i, j+1)
	}

	return cache[i][j]
}
