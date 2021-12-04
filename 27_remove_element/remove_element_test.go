package remove_element

import (
	"reflect"
	"testing"
)

type Test struct {
	in  []int
	val int
	out []int
	res int
}

var tests = []Test{
	{in: []int{3, 2, 2, 3}, val: 2, out: []int{3, 3, 2, 3}, res: 2},
	{in: []int{4, 2, 1, 4, 5, 4}, val: 4, out: []int{2, 1, 5, 4, 5, 4}, res: 3},
	{in: []int{0, 2, 0, 0}, val: 0, out: []int{2, 2, 0, 0}, res: 1},
	{in: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, out: []int{0, 1, 3, 0, 4, 0, 4, 2}, res: 5},
	{in: []int{}, val: 0, out: []int{}, res: 0},
}

func TestRemoveElement(t *testing.T) {
	for _, test := range tests {
		in := make([]int, len(test.in))
		copy(in, test.in)

		res := removeElement(in, test.val)

		if res != test.res || !reflect.DeepEqual(in, test.out) {
			t.Errorf("For %v and %d wanted %v and %d, got %v and %d\n", test.in, test.val, test.out, test.res, in, res)
		}
	}
}
