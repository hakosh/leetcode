package broken_calculator

import "testing"

type Test struct {
	startValue int
	target     int
	want       int
}

var tests = []Test{
	{2, 3, 2},
	{5, 8, 2},
	{3, 10, 3},
	{3, 17, 6},
	{10000004, 100006, 9899998},
}

func TestBrokenCalc(t *testing.T) {
	for _, test := range tests {
		if res := brokenCalc(test.startValue, test.target); res != test.want {
			t.Errorf("for %v and %v wanted %v, got %v", test.startValue, test.target, test.want, res)
		}
	}
}
