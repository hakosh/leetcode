package merge_sorted_array

import (
	"reflect"
	"testing"
)

type Test struct {
	nums1 []int
	n     int
	nums2 []int
	m     int
	out   []int
}

var tests = []Test{
	{nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3, out: []int{1, 2, 2, 3, 5, 6}},
	{nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{-2, 5, 6}, n: 3, out: []int{-2, 1, 2, 3, 5, 6}},
	{nums1: []int{1, 2, 3, 4}, m: 4, nums2: []int{}, n: 0, out: []int{1, 2, 3, 4}},
	{nums1: []int{0}, m: 0, nums2: []int{1}, n: 1, out: []int{1}},
}

func TestFindNumbers(t *testing.T) {
	for _, test := range tests {
		nums1 := make([]int, len(test.nums1))
		copy(nums1, test.nums1)

		merge(nums1, test.m, test.nums2, test.n)

		if !reflect.DeepEqual(nums1, test.out) {
			t.Errorf("For %v and %v wanted %v, got %v\n", test.nums1, test.nums2, test.out, nums1)
		}
	}
}
