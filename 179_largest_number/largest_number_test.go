package largest_number

import "testing"

type Test struct {
	nums []int
	out  string
}

var tests = []Test{
	{[]int{0, 0}, "0"},
	{[]int{10, 2}, "210"},
	{[]int{3, 30, 34, 5, 9}, "9534330"},
	{[]int{3, 30, 34, 310, 5, 9}, "9534331030"},
	{[]int{3, 30, 34, 310}, "34331030"},
	{[]int{30, 32, 34}, "343230"},
	{[]int{111311, 1113}, "1113111311"},
	{[]int{8308, 8308, 830}, "83088308830"},
}

func TestLargestNumber(t *testing.T) {
	for _, test := range tests {
		if res := largestNumber(test.nums); res != test.out {
			t.Errorf("For %v expected %q, got %q", test.nums, test.out, res)
		}
	}
}
