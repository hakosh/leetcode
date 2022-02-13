package maximum_value_after_insertion

import "testing"

type Test struct {
	n   string
	x   int
	out string
}

var tests = []Test{
	{"99", 9, "999"},
	{"468", 9, "9468"},
	{"468", 3, "4683"},
	{"412", 3, "4312"},
	{"-13", 2, "-123"},
	{"-568", 7, "-5678"},
	{"-648468153646", 5, "-5648468153646"},
}

func TestMaxValue(t *testing.T) {
	for _, test := range tests {
		if res := maxValue(test.n, test.x); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.n, test.x, test.out, res)
		}
	}
}
