package capacity_to_ship_packages_within_d_days

import "testing"

type Test struct {
	weights []int
	days    int
	out     int
}

var tests = []Test{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 15},
	{[]int{3, 2, 2, 4, 1, 4}, 3, 6},
	{[]int{1, 2, 3, 1, 1}, 4, 3},
}

func TestShipWithinDays(t *testing.T) {
	for _, test := range tests {
		if res := shipWithinDays(test.weights, test.days); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.weights, test.days, test.out, res)
		}
	}
}
