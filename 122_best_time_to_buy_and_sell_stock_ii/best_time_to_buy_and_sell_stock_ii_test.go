package best_time_to_buy_and_sell_stock_ii

import "testing"

type Test struct {
	in   []int
	want int
}

var cases = []Test{
	{in: []int{7, 1, 5, 3, 6, 4}, want: 7},
	{in: []int{7, 6, 4, 3, 1}, want: 0},
	{in: []int{1, 2, 3, 4, 5}, want: 4},
}

func TestMaxProfit(t *testing.T) {
	for _, test := range cases {
		if res := maxProfit(test.in); res != test.want {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.want, res)
		}
	}
}
