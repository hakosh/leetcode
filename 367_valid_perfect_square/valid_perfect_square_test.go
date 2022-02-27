package valid_perfect_square

import "testing"

type Test struct {
	num int
	out bool
}

var tests = []Test{
	{25, true},
	{16, true},
	{14, false},
	{4, true},
	{1, true},
}

func TestIsPerfectSquare(t *testing.T) {
	for _, test := range tests {
		if res := isPerfectSquare(test.num); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.num, test.out, res)
		}
	}
}
