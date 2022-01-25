package best_time_to_buy_and_sell_stock_iv

// BOTTOM UP

func maxProfit(k int, prices []int) int {
	dp := make([][][2]int, len(prices)+1)
	for i := range dp {
		dp[i] = make([][2]int, k+1)
	}

	for day := len(prices) - 1; day >= 0; day-- {
		for left := 1; left <= k; left++ {
			for holding := 0; holding < 2; holding++ {
				nothing := dp[day+1][left][holding]
				something := 0

				if holding == 1 {
					something = prices[day] + dp[day+1][left-1][0]
				} else {
					something = -prices[day] + dp[day+1][left][1]
				}

				dp[day][left][holding] = max(nothing, something)
			}
		}
	}

	return dp[0][k][0]
}

// TOP DOWN

var cache [][][2]int

func maxProfitR(k int, prices []int) int {
	cache = make([][][2]int, len(prices))
	for i := range cache {
		cache[i] = make([][2]int, k+1)
	}

	return dp(prices, 0, k, 0)
}

func dp(prices []int, day, left int, holding int) int {
	if day == len(prices) || left == 0 {
		return 0
	}

	if cache[day][left][holding] == 0 {
		profit := 0

		if holding == 1 {
			sell := prices[day] + dp(prices, day+1, left-1, 0)
			dont := dp(prices, day+1, left, 1)

			profit = max(sell, dont)
		} else {
			buy := -prices[day] + dp(prices, day+1, left, 1)
			not := dp(prices, day+1, left, 0)

			profit = max(buy, not)
		}

		cache[day][left][holding] = profit
	}

	return cache[day][left][holding]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
