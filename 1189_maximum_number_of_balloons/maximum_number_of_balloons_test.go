package maximum_number_of_balloons

import (
	"testing"
)

type Test struct {
	in  string
	out int
}

var tests = []Test{
	{"nlaebolko", 1},
	{"loonbalxballpoon", 2},
	{"hello", 0},
}

func TestDecodeString(t *testing.T) {
	for _, test := range tests {
		if res := maxNumberOfBalloons(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
