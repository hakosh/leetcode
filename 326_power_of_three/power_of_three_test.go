package power_of_three

import "testing"

type Test struct {
	in  int
	out bool
}

var tests = []Test{
	{9, true},
	{3, true},
	{1, true},
	{6, false},
	{129140164, false},
	{738622, false},
}

func TestIsPowerOfThree(t *testing.T) {
	for _, test := range tests {
		if res := isPowerOfThree(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}

func BenchmarkIsPowerOfThree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPowerOfThree(i)
	}
}
