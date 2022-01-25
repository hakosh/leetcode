package best_time_to_buy_and_sell_stock_with_cooldown

// BOTTOM UP

func maxProfit(prices []int) int {
	dp := make([][2]int, len(prices)+2)

	for day := len(prices) - 1; day >= 0; day-- {
		for holding := 0; holding < 2; holding++ {
			something := 0
			nothing := dp[day+1][holding]

			if holding == 1 {
				something = prices[day] + dp[day+2][0]
			} else {
				something = -prices[day] + dp[day+1][1]
			}

			dp[day][holding] = max(something, nothing)
		}
	}

	return dp[0][0]
}

// TOP DOWN

var cache [][2]int

func maxProfitR(prices []int) int {
	cache = make([][2]int, len(prices)+1)
	return dp(prices, 0, 0)
}

func dp(prices []int, day, holding int) int {
	if len(prices) <= day {
		return 0
	}

	if cache[day][holding] == 0 {
		something := 0
		nothing := dp(prices, day+1, holding)

		if holding == 1 {
			something = prices[day] + dp(prices, day+2, 0) // sell
		} else {
			something = -prices[day] + dp(prices, day+1, 1) // buy
		}

		cache[day][holding] = max(something, nothing)
	}

	return cache[day][holding]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
