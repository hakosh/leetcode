package two_sum_ii_input_array_is_sorted

import (
	"reflect"
	"testing"
)

type Test struct {
	numbers []int
	target  int
	out     []int
}

var tests = []Test{
	{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
	{[]int{2, 3, 4}, 6, []int{1, 3}},
	{[]int{-10, -4, 0, 1, 2, 3, 4}, -2, []int{2, 5}},
	{[]int{-1, 0}, -1, []int{1, 2}},
}

func TestTwoSum(t *testing.T) {
	for _, test := range tests {
		if res := twoSum(test.numbers, test.target); !reflect.DeepEqual(test.out, res) {
			t.Errorf("For %v and %v expected %v, got %v", test.numbers, test.target, test.out, res)
		}
	}
}
