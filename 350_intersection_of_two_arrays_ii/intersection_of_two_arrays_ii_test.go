package intersection_of_two_arrays_ii

import (
	"reflect"
	"testing"
)

type Test struct {
	nums1 []int
	nums2 []int
	want  []int
}

var cases = []Test{
	{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}, want: []int{2, 2}},
	{nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}, want: []int{4, 9}},
	{nums1: []int{4, 9, 5, 6}, nums2: []int{7, 8}, want: []int{}},
	{nums1: []int{1, 2}, nums2: []int{2, 1}, want: []int{1, 2}},
}

func TestIntersect(t *testing.T) {
	for _, test := range cases {
		res := intersect(test.nums1, test.nums2)

		if !reflect.DeepEqual(res, test.want) {
			t.Errorf("For %v and %v wanted %v, got %v\n", test.nums1, test.nums2, test.want, res)
		}
	}
}
