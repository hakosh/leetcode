package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	max := 0
	best := prices[0]

	for _, price := range prices {
		if price < best {
			best = price
		} else if price-best > max {
			max = price - best
		}
	}

	return max
}
