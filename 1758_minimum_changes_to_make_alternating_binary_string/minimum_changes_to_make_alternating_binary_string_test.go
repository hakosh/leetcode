package minimum_changes_to_make_alternating_binary_string

import "testing"

type Test struct {
	s   string
	out int
}

var tests = []Test{
	{"0100", 1},
	{"10", 0},
	{"1111", 2},
	{"10010", 2},
	{"10010100", 3},
}

func TestMinOperations(t *testing.T) {
	for _, test := range tests {
		if res := minOperations(test.s); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.s, test.out, res)
		}
	}
}
