package house_robber

import (
	"testing"
)

type Test struct {
	in  []int
	out int
}

var cases = []Test{
	{in: []int{1, 2, 3, 1}, out: 4},
	{in: []int{2, 7, 9, 3, 1}, out: 12},
	{in: []int{1, 2}, out: 2},
	{in: []int{1, 2, 4}, out: 5},
	{in: []int{0}, out: 0},
	{in: []int{0, 0}, out: 0},
}

func TestRob(t *testing.T) {
	for _, test := range cases {
		res := rob(test.in)

		if res != test.out {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.out, res)
		}
	}
}
