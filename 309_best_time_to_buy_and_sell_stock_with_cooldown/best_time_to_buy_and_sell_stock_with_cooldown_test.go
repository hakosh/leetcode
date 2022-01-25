package best_time_to_buy_and_sell_stock_with_cooldown

import "testing"

type Test struct {
	prices []int
	out    int
}

var tests = []Test{
	{[]int{1, 2, 3, 0, 2}, 3},
	//{[]int{1}, 0},
}

func TestMaxProfit(t *testing.T) {
	for _, test := range tests {
		if res := maxProfit(test.prices); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.prices, test.out, res)
		}
	}
}
