package rotate_string

import "testing"

type Test struct {
	s    string
	goal string
	out  bool
}

var tests = []Test{
	{"leetcode", "etcodele", true},
}

func TestRotateString(t *testing.T) {
	for _, test := range tests {
		if res := rotateString(test.s, test.goal); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.s, test.goal, test.out, res)
		}
	}
}
