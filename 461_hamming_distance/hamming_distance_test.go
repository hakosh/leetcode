package hamming_distance

import "testing"

type Test struct {
	x   int
	y   int
	out int
}

var tests = []Test{
	{1, 4, 2},
	{3, 1, 1},
	{0, 1, 1},
	{0, 0, 0},
	{5, 5, 0},
}

func TestHammingDistance(t *testing.T) {
	for _, test := range tests {
		if res := hammingDistance(test.x, test.y); res != test.out {
			t.Errorf("For %b and %b expected %d, got %d", test.x, test.y, test.out, res)
		}
	}
}
