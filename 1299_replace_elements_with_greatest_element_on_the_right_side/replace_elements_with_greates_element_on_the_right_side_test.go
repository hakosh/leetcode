package replace_elements_with_greatest_element_on_the_right_side

import (
	"reflect"
	"testing"
)

type Test struct {
	in  []int
	out []int
}

var tests = []Test{
	{in: []int{17, 18, 5, 4, 6, 1}, out: []int{18, 6, 6, 6, 1, -1}},
	{in: []int{1}, out: []int{-1}},
	{in: []int{0, 0, 0}, out: []int{0, 0, -1}},
}

func TestReplaceElements(t *testing.T) {
	for _, test := range tests {
		res := make([]int, len(test.in))
		copy(res, test.in)

		replaceElements(res)

		if !reflect.DeepEqual(res, test.out) {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.out, res)
		}
	}
}
