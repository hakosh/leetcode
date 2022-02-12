package subarray_sum_equals_k

import "testing"

type Test struct {
	nums []int
	k    int
	out  int
}

var tests = []Test{
	{[]int{1, 1, 1}, 2, 2},
	{[]int{1, 2, 3}, 3, 2},
	{[]int{2, 2, -2, 2}, 4, 2},
}

func TestSubarraySum(t *testing.T) {
	for _, test := range tests {
		if res := subarraySum(test.nums, test.k); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.nums, test.k, test.out, res)
		}
	}
}
