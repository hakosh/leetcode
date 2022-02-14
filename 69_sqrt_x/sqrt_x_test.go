package sqrt_x

import "testing"

type Test struct {
	x   int
	out int
}

var tests = []Test{
	{4, 2},
	{8, 2},
	{1, 1},
	{64, 8},
	{65, 8},
}

func TestMySqrt(t *testing.T) {
	for _, test := range tests {
		if res := mySqrt(test.x); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.x, test.out, res)
		}
	}
}
