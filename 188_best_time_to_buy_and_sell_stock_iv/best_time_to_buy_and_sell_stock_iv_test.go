package best_time_to_buy_and_sell_stock_iv

import "testing"

type Test struct {
	k      int
	prices []int
	out    int
}

var tests = []Test{
	{2, []int{2, 4, 1}, 2},
	{0, []int{2, 4, 1}, 0},
	{2, []int{3, 2, 6, 5, 0, 3}, 7},
}

func TestMaxProfit(t *testing.T) {
	for _, test := range tests {
		if res := maxProfit(test.k, test.prices); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.k, test.prices, test.out, res)
		}
	}
}
