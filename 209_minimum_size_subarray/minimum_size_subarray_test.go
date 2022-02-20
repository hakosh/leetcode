package minimum_size_subarray

import "testing"

type Test struct {
	target int
	nums   []int
	out    int
}

var tests = []Test{
	{7, []int{2, 3, 1, 2, 4, 3}, 2},
	{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
	{4, []int{1, 4, 4}, 1},
	{4, []int{4}, 1},
	{11, []int{1, 2, 3, 4, 5}, 3},
	{15, []int{1, 2, 3, 4, 5}, 5},
}

func TestMinSubarrayLen(t *testing.T) {
	for _, test := range tests {
		if res := minSubArrayLen(test.target, test.nums); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.target, test.nums, test.out, res)
		}
	}
}
