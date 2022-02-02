package missing_number

import (
	"testing"
)

type Test struct {
	in  []int
	out int
}

var tests = []Test{
	{[]int{0, 1, 2, 4, 5, 6}, 3},
	{[]int{1, 2, 3, 4, 5, 6}, 0},
	{[]int{1}, 0},
	{[]int{0, 1}, 2},
}

func TestMissingNumber(t *testing.T) {
	for _, test := range tests {
		if res := missingNumber(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
