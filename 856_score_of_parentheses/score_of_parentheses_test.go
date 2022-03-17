package score_of_parentheses

import "testing"

type Test struct {
	s    string
	want int
}

var tests = []Test{
	{"()", 1},
	{"(())", 2},
	{"()()", 2},
	{"((()))()", 5},
	{"((()()))", 8},
}

func TestScoreOfParentheses(t *testing.T) {
	for _, test := range tests {
		if res := scoreOfParentheses(test.s); res != test.want {
			t.Errorf("for %q wanted %v, got %v", test.s, test.want, res)
		}
	}
}
