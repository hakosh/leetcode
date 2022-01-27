package maximum_sum_circular_subarray

import "testing"

type Test struct {
	in  []int
	out int
}

var tests = []Test{
	{[]int{1, -2, 3, -2}, 3},
	{[]int{5, -3, 5}, 10},
	{[]int{-3, -2, -3}, -2},
	{[]int{3, -1, 2, -1}, 4},
	{[]int{3, -2, 2, -3}, 3},
	{[]int{-5, 3, 5}, 8},
}

func TestMaxSubarraySumCircular(t *testing.T) {
	for _, test := range tests {
		if res := maxSubarraySumCircular(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
