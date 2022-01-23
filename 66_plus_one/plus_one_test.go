package plus_one

import (
	"reflect"
	"testing"
)

type Test struct {
	in  []int
	out []int
}

var tests = []Test{
	{[]int{1, 2, 3}, []int{1, 2, 4}},
	{[]int{9}, []int{1, 0}},
	{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
	{[]int{9, 9, 9, 9}, []int{1, 0, 0, 0, 0}},
}

func TestPlusOne(t *testing.T) {
	for _, test := range tests {
		if res := plusOne(test.in); !reflect.DeepEqual(res, test.out) {
			t.Errorf("For %v wanted %v, got %v", test.in, test.out, res)
		}
	}
}
