package sort_array_by_parity

import (
	"reflect"
	"testing"
)

type Test struct {
	in   []int
	want []int
}

var cases = []Test{
	{in: []int{3, 1, 2, 4}, want: []int{2, 4, 3, 1}},
	{in: []int{1}, want: []int{1}},
	{in: []int{1, 0, 3}, want: []int{0, 1, 3}},
}

func TestMoveZeroes(t *testing.T) {
	for _, test := range cases {
		in := make([]int, len(test.in))
		copy(in, test.in)
		in = sortArrayByParity(in)

		if !reflect.DeepEqual(in, test.want) {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.want, in)
		}
	}
}
