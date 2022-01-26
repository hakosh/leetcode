package coin_change_2

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for j := range coins {
		for i := coins[j]; i < amount+1; i++ {
			dp[i] += dp[i-coins[j]]
		}
	}

	return dp[amount]
}
