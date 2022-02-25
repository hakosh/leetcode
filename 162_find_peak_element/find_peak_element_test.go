package find_peak_element

import "testing"

type Test struct {
	nums []int
	out  int
}

var tests = []Test{
	{[]int{1, 2, 3, 1}, 2},
	{[]int{1, 2, 1, 3, 5, 6, 4}, 5},
	{[]int{8, 9, 8, 7, 6, 5, 4, 3}, 1},
	{[]int{1}, 0},
}

func TestFindPeakElements(t *testing.T) {
	for _, test := range tests {
		if res := findPeakElement(test.nums); test.out != res {
			t.Errorf("For %v expected %v, got %v", test.nums, test.out, res)
		}
	}
}
