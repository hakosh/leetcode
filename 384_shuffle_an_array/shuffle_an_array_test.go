package shuffle_an_array

import (
	"math/rand"
	"reflect"
	"testing"
)

type Test struct {
	in  []int
	out []int
}

var tests = []Test{
	{[]int{1, 2, 3}, []int{2, 3, 1}},
	{[]int{4, 5, 6, 7, 8, 9}, []int{5, 8, 4, 6, 7, 9}},
}

func TestSolution(t *testing.T) {
	for _, test := range tests {
		rand.Seed(1)
		s := Constructor(test.in)

		if res := s.Shuffle(); !reflect.DeepEqual(res, test.out) {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}

func TestSolution_Reset(t *testing.T) {
	nums := tests[1].in

	s := Constructor(nums)

	if reflect.DeepEqual(nums, s.Shuffle()) {
		t.Errorf("Array was not shuffled")
	}

	if !reflect.DeepEqual(nums, s.Reset()) {
		t.Errorf("Array was not reset")
	}
}
