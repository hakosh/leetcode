package best_time_to_buy_and_sell_stock_with_transaction_fee

func maxProfit(prices []int, fee int) int {
	dp := make([][2]int, len(prices)+1)

	for day := len(prices) - 1; day >= 0; day-- {
		for holding := 0; holding < 2; holding++ {
			something := 0
			nothing := dp[day+1][holding]

			if holding == 1 {
				something = prices[day] - fee + dp[day+1][0]
			} else {
				something = -prices[day] + dp[day+1][1]
			}

			dp[day][holding] = max(something, nothing)
		}
	}

	return dp[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
