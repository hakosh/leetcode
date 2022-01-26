package paint_fence

func numWays(n int, k int) int {
	dp := make([]int, n+1)

	dp[0] = k
	dp[1] = k * k

	for i := 2; i < n; i++ {
		same := 1 * (k - 1) * dp[i-2]
		diff := (k - 1) * dp[i-1]

		dp[i] = same + diff
	}

	return dp[n-1]
}
