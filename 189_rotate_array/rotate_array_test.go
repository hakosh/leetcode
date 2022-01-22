package rotate_array

import (
	"reflect"
	"testing"
)

type Test struct {
	in  []int
	k   int
	out []int
}

var tests = []Test{
	{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
	{[]int{1, 2, 3, 4, 5, 6, 7}, 5, []int{3, 4, 5, 6, 7, 1, 2}},
	{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
}

func TestRotate(t *testing.T) {
	for _, test := range tests {
		in := make([]int, len(test.in))
		copy(in, test.in)

		rotate(in, test.k)

		if !reflect.DeepEqual(in, test.out) {
			t.Errorf("For %v and %v wanted %v, got %v", test.in, test.k, test.out, in)
		}
	}
}
