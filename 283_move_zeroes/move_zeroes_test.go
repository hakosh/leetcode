package move_zeroes

import (
	"reflect"
	"testing"
)

type Test struct {
	in []int
	want []int
}

var cases = []Test{
	{ in: []int{}, want: []int{} },
	{ in: []int{0, 1, 0, 3, 12}, want: []int{1, 3, 12, 0, 0} },
	{ in: []int{1, 0, 0, 0, 9, 0}, want: []int{1, 9, 0, 0, 0, 0}},
}

func TestMoveZeroes(t *testing.T) {
	for _, test := range cases {
		in := make([]int, len(test.in))
		copy(in, test.in)
		moveZeroes(in)

		if !reflect.DeepEqual(in, test.want) {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.want, in)
		}
	}
}
