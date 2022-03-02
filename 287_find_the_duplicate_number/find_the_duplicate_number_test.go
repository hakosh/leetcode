package find_the_duplicate_number

import "testing"

type Test struct {
	nums []int
	out  int
}

var tests = []Test{
	{[]int{1, 3, 4, 2, 2}, 2},
	{[]int{3, 1, 3, 4, 2}, 3},
	{[]int{1, 1, 2}, 1},
	{[]int{1, 2, 3, 4, 5, 5}, 5},
	{[]int{2, 2, 2, 2, 2}, 2},
}

func TestFindDuplicate(t *testing.T) {
	for _, test := range tests {
		if res := findDuplicate(test.nums); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.nums, test.out, res)
		}
	}
}
