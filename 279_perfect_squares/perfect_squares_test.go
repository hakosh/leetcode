package perfect_squares

import "testing"

type Test struct {
	in  int
	out int
}

var tests = []Test{
	{12, 3},
	{13, 2},
	{16, 1},
}

func TestNumSquares(t *testing.T) {
	for _, test := range tests {
		if res := numSquares(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
