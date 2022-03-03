package arithmetic_slices

import "testing"

type Test struct {
	nums []int
	out  int
}

var tests = []Test{
	{[]int{1, 2, 3, 4}, 3},
	{[]int{1, 3, 5, 7, 9}, 6},
	{[]int{1, 3, 5, 4, 9, 11, 13}, 2},
	{[]int{1}, 0},
}

func TestNumberOfArithmeticSlices(t *testing.T) {
	for _, test := range tests {
		if res := numberOfArithmeticSlices(test.nums); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.nums, test.out, res)
		}
	}
}
