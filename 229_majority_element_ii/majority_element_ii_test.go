package majority_element_ii

import (
	"reflect"
	"testing"
)

type Test struct {
	nums []int
	out  []int
}

var tests = []Test{
	{[]int{3, 2, 3}, []int{3}},
	{[]int{1}, []int{1}},
	{[]int{1, 2}, []int{1, 2}},
	{[]int{1, 2, 3, 2, 2, 1}, []int{2}},
}

func TestMajorityElement(t *testing.T) {
	for _, test := range tests {
		if res := majorityElement(test.nums); !reflect.DeepEqual(test.out, res) {
			t.Errorf("For %v expected %v, got %v", test.nums, test.out, res)
		}
	}
}
