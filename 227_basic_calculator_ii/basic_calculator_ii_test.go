package basic_calculator_ii

import "testing"

type Test struct {
	expr string
	out  int
}

var tests = []Test{
	{"", 0},
	{"3/2", 1},
	{"2+4/2", 4},
	{"359+451/96", 363},
	{"3+2*2-5", 2},
	{"3+2*2-5/8+24-9/3*7+1", 11},
}

func TestCalculate(t *testing.T) {
	for _, test := range tests {
		if res := calculate(test.expr); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.expr, test.out, res)
		}
	}
}
