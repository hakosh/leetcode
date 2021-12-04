package check_if_n_and_its_double_exist

import (
	"testing"
)

type Test struct {
	in  []int
	out bool
}

var tests = []Test{
	{in: []int{1, 1, 2}, out: true},
	{in: []int{0, 1, 4, 4, 2}, out: true},
	{in: []int{10, 2, 5, 3}, out: true},
	{in: []int{7, 1, 14, 11}, out: true},
	{in: []int{3, 1, 7, 11}, out: false},
	{in: []int{-5, -5}, out: false},
	{in: []int{-2, 0, 10, -19, 4, 6, -8}, out: false},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, test := range tests {
		in := make([]int, len(test.in))
		copy(in, test.in)

		if res := checkIfExist(test.in); res != test.out {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.out, res)
		}
	}
}
