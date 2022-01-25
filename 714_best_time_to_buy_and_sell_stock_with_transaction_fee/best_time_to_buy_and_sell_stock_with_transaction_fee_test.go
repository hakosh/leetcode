package best_time_to_buy_and_sell_stock_with_transaction_fee

import "testing"

type Test struct {
	in   []int
	fee  int
	want int
}

var cases = []Test{
	{in: []int{1, 2, 3, 4, 5}, fee: 2, want: 2},
	{in: []int{1, 3, 2, 8, 4, 9}, fee: 2, want: 8},
	{in: []int{1, 3, 7, 5, 10, 3}, fee: 3, want: 6},
}

func TestMaxProfit(t *testing.T) {
	for _, test := range cases {
		if res := maxProfit(test.in, test.fee); res != test.want {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.want, res)
		}
	}
}
