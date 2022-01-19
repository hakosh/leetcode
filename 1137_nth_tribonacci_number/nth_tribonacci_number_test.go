package nth_tribonacci_number

import "testing"

type Test struct {
	in  int
	out int
}

var tests = []Test{
	{0, 0},
	{4, 4},
}

func TestDecodeString(t *testing.T) {
	for _, test := range tests {
		if res := tribonacci(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
