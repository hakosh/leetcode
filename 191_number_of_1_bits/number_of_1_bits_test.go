package number_of_1_bits

import "testing"

type Test struct {
	n   uint32
	out int
}

var tests = []Test{
	{11, 4},
	{0, 0},
}

func TestHammingWeight(t *testing.T) {
	for _, test := range tests {
		if res := hammingWeight(test.n); res != test.out {
			t.Errorf("For %b expected %v, got %v", test.n, test.out, res)
		}
	}
}
