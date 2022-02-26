package find_k_closest_elements

import (
	"reflect"
	"testing"
)

type Test struct {
	arr []int
	k   int
	x   int
	out []int
}

var tests = []Test{
	{[]int{1, 2, 3, 4, 5}, 4, 3, []int{1, 2, 3, 4}},
	{[]int{1, 2, 3, 4, 5}, 4, -1, []int{1, 2, 3, 4}},
	{[]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5, []int{3, 3, 4}},
	{[]int{3, 5, 8, 10}, 2, 15, []int{8, 10}},
	{[]int{1, 3}, 1, 2, []int{1}},
}

func TestFindClosestElements(t *testing.T) {
	for _, test := range tests {
		if res := findClosestElements(test.arr, test.k, test.x); !reflect.DeepEqual(res, test.out) {
			t.Errorf("For %v, %v, and %v expected %v, got %v", test.arr, test.k, test.x, test.out, res)
		}
	}
}
