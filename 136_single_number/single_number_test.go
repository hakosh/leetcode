package single_number

import (
	"testing"
)

type Test struct {
	in  []int
	out int
}

var tests = []Test{
	{[]int{2, 2, 1}, 1},
	{[]int{4, 1, 2, 1, 2}, 4},
	{[]int{1, 2, 3, 4, 5, 1, 2, 3, 4}, 5},
}

func TestRotate(t *testing.T) {
	for _, test := range tests {
		if res := singleNumber(test.in); res != test.out {
			t.Errorf("For %v wanted %v, got %v", test.in, test.out, res)
		}
	}
}
