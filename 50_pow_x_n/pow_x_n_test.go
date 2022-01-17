package pow_x_n

import (
	"testing"
)

type Test struct {
	x   float64
	n   int
	out float64
}

var tests = []Test{
	{2.0, 2, 4.0},
	{2.0, 10, 1024.0},
	{2.1, 3, 9.261},
	{2.0, -2, 0.25},
	{0.00001, 2147483647, 0},
	{1, -2147483647, 1},
}

func TestDecodeString(t *testing.T) {
	for _, test := range tests {
		if res := myPow(test.x, test.n); res != test.out {
			t.Errorf("For %v^%v expected %v, got %v", test.x, test.n, test.out, res)
		}
	}
}
