package valid_parentheses

import "testing"

type Test struct {
	in  string
	out bool
}

var tests = []Test{
	{"()", true},
	{"(){}[]", true},
	{"({})", true},
	{"[(])", false},
	{"(}", false},
}

func TestIsValid(t *testing.T) {
	for _, test := range tests {
		if res := isValid(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
