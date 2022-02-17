package add_binary

import "testing"

type Test struct {
	a   string
	b   string
	out string
}

var tests = []Test{
	{"11", "1", "100"},
	{"1010", "1011", "10101"},
}

func TestAddBinary(t *testing.T) {
	for _, test := range tests {
		if res := addBinary(test.a, test.b); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.a, test.b, test.out, res)
		}
	}
}
