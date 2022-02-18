package array_partition_i

import "testing"

type Test struct {
	nums []int
	out  int
}

var tests = []Test{
	{[]int{6, 2, 6, 5, 1, 2}, 9},
}

func TestArrayPairSum(t *testing.T) {
	for _, test := range tests {
		if res := arrayPairSum(test.nums); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.nums, test.out, res)
		}
	}
}
