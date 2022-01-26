package coin_change_2

import "testing"

type Test struct {
	amount int
	coins  []int
	out    int
}

var tests = []Test{
	{2, []int{2}, 1},
	{5, []int{1, 2, 5}, 4},
	{10, []int{5}, 1},
	{500, []int{2, 7, 13}, 717},
	{35, []int{1, 2, 5, 10}, 140},
}

func TestChange(t *testing.T) {
	for _, test := range tests {
		if res := change(test.amount, test.coins); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.amount, test.coins, test.out, res)
		}
	}
}
