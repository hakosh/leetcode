package remove_duplicates

import (
	"reflect"
	"testing"
)

type Test struct {
	in  []int
	out []int
	res int
}

var tests = []Test{
	{in: []int{1, 1, 2}, out: []int{1, 2}, res: 2},
	{in: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, out: []int{0, 1, 2, 3, 4}, res: 5},
	{in: []int{1}, out: []int{1}, res: 1},
	{in: []int{1, 2}, out: []int{1, 2}, res: 2},
	{in: []int{1, 2, 3}, out: []int{1, 2, 3}, res: 3},
	{in: []int{}, out: []int{}, res: 0},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, test := range tests {
		in := make([]int, len(test.in))
		copy(in, test.in)

		res := removeDuplicates(in)

		if res != test.res || !reflect.DeepEqual(in[:res], test.out) {
			t.Errorf("For %v wanted %v and %d, got %v and %d\n", test.in, test.out, test.res, in[:res], res)
		}
	}
}
