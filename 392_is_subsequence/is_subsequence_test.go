package is_subsequence

import "testing"

type Test struct {
	s   string
	t   string
	out bool
}

var tests = []Test{
	{"", "", true},
	{"abc", "ahbgdc", true},
	{"axc", "ahbgdc", false},
	{"axc", "", false},
	{"", "majom", true},
}

func TestIsSubsequence(t *testing.T) {
	for _, test := range tests {
		if res := isSubsequence(test.s, test.t); res != test.out {
			t.Errorf("For %q and %q expected %v, got %v", test.s, test.t, test.out, res)
		}
	}
}
