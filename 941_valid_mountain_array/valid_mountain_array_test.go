package valid_mountain_array

import (
	"testing"
)

type Test struct {
	in  []int
	out bool
}

var tests = []Test{
	{in: []int{1, 2}, out: false},
	{in: []int{3, 5, 5}, out: false},
	{in: []int{0, 3, 2, 1}, out: true},
	{in: []int{4, 3, 2, 1}, out: false},
	{in: []int{1, 2, 3, 4}, out: false},
	{in: []int{1, 1, 1, 1}, out: false},
	{in: []int{1, 1, 2, 2, 2, 1, 1}, out: false},
	{in: []int{0, 1, 2, 4, 2, 1}, out: true},
}

func TestValidMountainArray(t *testing.T) {
	for _, test := range tests {
		in := make([]int, len(test.in))
		copy(in, test.in)

		if res := validMountainArray(test.in); res != test.out {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.out, res)
		}
	}
}
