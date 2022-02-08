package find_first_and_last_position_of_element_in_sorted_array

import (
	"reflect"
	"testing"
)

type Test struct {
	in     []int
	target int
	out    []int
}

var tests = []Test{
	{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
	{[]int{5, 7, 7, 8, 8, 10}, 10, []int{5, 5}},
	{[]int{5, 7, 7, 8, 9, 10}, 9, []int{4, 4}},
	{[]int{1, 2, 2, 2, 2, 2, 2, 3}, 2, []int{1, 6}},
	{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
	{[]int{1, 2}, 1, []int{0, 0}},
	{[]int{2, 2}, 2, []int{0, 1}},
	{[]int{2, 2}, 3, []int{-1, -1}},
	{[]int{3, 3, 3, 3, 3}, 3, []int{0, 4}},
	{[]int{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3}, 2, []int{4, 7}},
	{[]int{}, 2, []int{-1, -1}},
}

func TestSearchRange(t *testing.T) {
	for _, test := range tests {
		if res := searchRange(test.in, test.target); !reflect.DeepEqual(test.out, res) {
			t.Errorf("For %v and %v expected %v, got %v", test.in, test.target, test.out, res)
		}
	}
}
