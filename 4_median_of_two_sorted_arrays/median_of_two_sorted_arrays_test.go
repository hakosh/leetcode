package median_of_two_sorted_arrays

import "testing"

type Test struct {
	nums1 []int
	nums2 []int
	out   float64
}

var tests = []Test{
	{[]int{}, []int{}, 0},
	{[]int{}, []int{9}, 9},
	{[]int{3}, []int{9}, 6},
	{[]int{1, 3, 5, 7}, []int{9}, 5},
	{[]int{1, 2}, []int{3, 4}, 2.5},
	{[]int{3, 4}, []int{1, 2}, 2.5},
	{[]int{1, 2, 3, 4, 5}, []int{3, 6, 7, 9, 12}, 4.5},
	{[]int{1, 2, 3, 4, 5, 6, 8, 11, 12}, []int{3, 6, 7}, 5.5},
	{[]int{2, 2, 4, 4}, []int{2, 2, 4, 4}, 3},
	{[]int{0, 0, 0, 0, 0}, []int{-1, 0, 0, 0, 0, 0, 1}, 0},
	{[]int{}, []int{1, 2, 3, 4, 5, 6}, 3.5},
	{[]int{3, 5, 6}, []int{1, 2, 4}, 3.5},
	{[]int{4, 5, 6}, []int{1, 2, 3}, 3.5},
	{[]int{3}, []int{1, 2, 4, 5}, 3},
}

func TestFindMedianSortedArrays(t *testing.T) {
	for _, test := range tests {
		if res := findMedianSortedArrays(test.nums1, test.nums2); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.nums1, test.nums2, test.out, res)
		}
	}
}
