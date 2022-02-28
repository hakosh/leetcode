package find_minimum_in_rotated_sorted_array_ii

import "testing"

type Test struct {
	nums []int
	out  int
}

var tests = []Test{
	{[]int{2, 2, 2, 0, 1}, 0},
	{[]int{0, 1, 4, 4, 5, 6, 7}, 0},
	{[]int{4, 5, 6, 7, 0, 1, 4}, 0},
	{[]int{5, 5, 5, 5, 5}, 5},
	{[]int{10, 1, 10, 10, 10}, 1},
	{[]int{10, 10, 10, 1, 10}, 1},
	{[]int{1, 1, 3, 1}, 1},
	{[]int{1, 1, 1}, 1},
}

func TestFindMin(t *testing.T) {
	for _, test := range tests {
		if res := findMin(test.nums); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.nums, test.out, res)
		}
	}
}
