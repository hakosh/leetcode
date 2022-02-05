package backspace_string_compare

import "testing"

type Test struct {
	s   string
	t   string
	out bool
}

var tests = []Test{
	{"a##c", "#a#c", true},
	{"a", "a", true},
	{"a", "b", false},
	{"a", "bca", false},
	{"", "", true},
	{"a#", "b#", true},
	{"a#c", "b", false},
	{"ab##", "c#d#", true},
	{"abc##x", "a #x", true},
	{"bxj##tw", "bxo#j##tw", true},
	{"xywrrmp", "xywrrmu#p", true},
}

func TestBackspaceCompare(t *testing.T) {
	for _, test := range tests {
		if res := backspaceCompare(test.s, test.t); res != test.out {
			t.Errorf("For %q and %q expected %v, got %v", test.s, test.t, test.out, res)
		}
	}
}
