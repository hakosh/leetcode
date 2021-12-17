package best_time_to_buy_and_sell_stock

import (
	"testing"
)

type Test struct {
	in   []int
	want int
}

var cases = []Test{
	{in: []int{7, 1, 5, 3, 6, 4}, want: 5},
	{in: []int{7, 6, 4, 3, 1}, want: 0},
}

func TestThirdMax(t *testing.T) {
	for _, test := range cases {
		if res := maxProfit(test.in); res != test.want {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.want, res)
		}
	}
}
