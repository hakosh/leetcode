package guess_number_higher_or_lower

import "testing"

type Test struct {
	n   int
	num int
}

var tests = []Test{
	{10, 6},
	{1, 1},
	{2, 1},
	{979, 415},
}

func TestGuessNumber(t *testing.T) {
	for _, test := range tests {
		setNum(test.num)

		if res := guessNumber(test.n); res != test.num {
			t.Errorf("For %v and %v expected %v, got %v", test.n, test.num, test.num, res)
		}
	}
}
