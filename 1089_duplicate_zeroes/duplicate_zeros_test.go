package duplicate_zeros

import (
	"reflect"
	"testing"
)

type Test struct {
	in  []int
	out []int
}

var tests = []Test{
	{in: []int{1, 0, 2, 3, 0, 4, 5, 0}, out: []int{1, 0, 0, 2, 3, 0, 0, 4}},
	{in: []int{1, 2, 3}, out: []int{1, 2, 3}},
	{in: []int{1, 2, 0, 0}, out: []int{1, 2, 0, 0}},
	{in: []int{0, 0, 5, 0, 0}, out: []int{0, 0, 0, 0, 5}},
	{in: []int{0, 0, 0, 0, 0}, out: []int{0, 0, 0, 0, 0}},
}

func TestFindNumbers(t *testing.T) {
	for _, test := range tests {
		res := make([]int, len(test.in))
		copy(res, test.in)

		duplicateZeros(res)

		if !reflect.DeepEqual(res, test.out) {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.out, res)
		}
	}
}
