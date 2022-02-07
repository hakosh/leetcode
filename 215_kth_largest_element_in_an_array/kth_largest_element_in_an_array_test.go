package kth_largest_element_in_an_array

import (
	"testing"
)

type Test struct {
	in  []int
	k   int
	out int
}

var tests = []Test{
	{[]int{3, 2, 1, 5, 6, 4}, 6, 1},
	{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	{[]int{3}, 1, 3},
	{[]int{-1, -2, -3}, 1, -1},
	{[]int{1, 2, 3, 4, 5}, 3, 3},
	{[]int{26, 45, 64, 18, 69, 54, 26, 43, 24, 99, 17, 35, 83, 94, 69, 72, 13, 35, 19, 70, 14, 20, 34, 13, 81, 54, 59, 27, 19, 15, 17, 2, 26, 78, 45, 34, 62, 18, 55, 82, 2, 93, 1, 55, 97, 37, 48, 28, 20, 72, 61, 10, 54, 12, 5, 3, 19, 31, 86, 12, 4, 7, 97, 97, 20, 24, 2, 14, 92, 94, 82, 97, 15, 30, 76, 57, 5, 69, 27, 59, 91, 19, 93, 69, 60, 1, 27, 74, 62, 91, 79, 31, 33, 82, 86, 14, 42, 84, 26, 52}, 60, 28},
}

func TestFindKthLargest(t *testing.T) {
	for _, test := range tests {
		if res := findKthLargest(test.in, test.k); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.in, test.k, test.out, res)
		}
	}
}

func BenchmarkFindKthLargest(b *testing.B) {
	test := tests[len(tests)-1]
	for i := 0; i < b.N; i++ {
		findKthLargest(test.in, test.k)
	}
}
