package decode_ways

import "testing"

type Test struct {
	in  string
	out int
}

var tests = []Test{
	{"12", 2},
	{"226", 3},
	{"06", 0},
	{"370", 0},
	{"210", 1},
	{"2324", 4},
}

func TestNumDecodings(t *testing.T) {
	for _, test := range tests {
		if res := numDecodings(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
