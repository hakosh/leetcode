package count_all_valid_pickup_and_delivery_options

import "testing"

type Test struct {
	n   int
	out int
}

var tests = []Test{
	{1, 1},
	{2, 6},
	{3, 90},
}

func TestCountOrders(t *testing.T) {
	for _, test := range tests {
		if res := countOrders(test.n); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.n, test.out, res)
		}
	}
}
