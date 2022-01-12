package evaluate_reverse_polish_notation

import (
	"testing"
)

type Test struct {
	in  []string
	out int
}

var tests = []Test{
	{[]string{"2", "1", "+", "3", "*"}, 9},
	{[]string{"4", "13", "5", "/", "+"}, 6},
	{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	{[]string{"2"}, 2},
}

func TestDailyTemperatures(t *testing.T) {
	for _, test := range tests {
		if res := evalRPN(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
