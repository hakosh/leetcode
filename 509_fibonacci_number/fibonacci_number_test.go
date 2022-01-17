package fibonacci_number

import (
	"testing"
)

type Test struct {
	in  int
	out int
}

var tests = []Test{
	{0, 0},
	{1, 1},
	{3, 2},
	{4, 3},
}

func TestDecodeString(t *testing.T) {
	for _, test := range tests {
		if res := fib(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, test.in)
		}
	}
}
