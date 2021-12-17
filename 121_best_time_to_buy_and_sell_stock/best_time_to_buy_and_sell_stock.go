package best_time_to_buy_and_sell_stock

import "math"

func maxProfit(prices []int) int {
	best := math.MaxInt32
	max := 0

	for _, buy := range prices {
		if buy < best {
			best = buy
		} else if buy-best > max {
			max = buy - best
		}
	}

	return max
}
