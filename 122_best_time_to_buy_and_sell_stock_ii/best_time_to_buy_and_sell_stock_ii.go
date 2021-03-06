package best_time_to_buy_and_sell_stock_ii

func maxProfit(prices []int) int {
	total := 0

	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			total += prices[i+1] - prices[i]
		}
	}

	return total
}
