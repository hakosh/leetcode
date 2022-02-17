package nim_game

import "testing"

type Test struct {
	n   int
	out bool
}

var tests = []Test{
	{4, false},
	{5, true},
}

func TestCanWinNim(t *testing.T) {
	for _, test := range tests {
		if res := canWinNim(test.n); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.n, test.out, res)
		}
	}
}
