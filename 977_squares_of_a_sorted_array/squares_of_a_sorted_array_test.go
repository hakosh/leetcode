package squares_of_a_sorted_array

import (
	"reflect"
	"testing"
)

type Test struct {
	in  []int
	out []int
}

var tests = []Test{
	{in: []int{-4, -1, 0, 3, 10}, out: []int{0, 1, 9, 16, 100}},
	{in: []int{-7, -3, 2, 3, 11}, out: []int{4, 9, 9, 49, 121}},
	{in: []int{-12, -9, 2, 3, 8}, out: []int{4, 9, 64, 81, 144}},
	{in: []int{-5, -3, -2, -1}, out: []int{1, 4, 9, 25}},
	{in: []int{1}, out: []int{1}},
}

func TestSortedSquares(t *testing.T) {
	for _, test := range tests {
		res := sortedSquares(test.in)

		if !reflect.DeepEqual(res, test.out) {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.out, res)
		}
	}
}
