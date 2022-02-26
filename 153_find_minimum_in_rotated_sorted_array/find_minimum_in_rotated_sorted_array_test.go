package find_minimum_in_rotated_sorted_array

import "testing"

type Test struct {
	nums []int
	out  int
}

var tests = []Test{
	{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
	{[]int{3, 4, 5, 1, 2}, 1},
	{[]int{11, 13, 15, 17}, 11},
}

func TestFindMin(t *testing.T) {
	for _, test := range tests {
		if res := findMin(test.nums); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.nums, test.out, res)
		}
	}
}
