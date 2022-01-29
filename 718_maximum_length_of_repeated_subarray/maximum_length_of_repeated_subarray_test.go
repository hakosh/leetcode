package maximum_length_of_repeated_subarray

import "testing"

type Test struct {
	nums1 []int
	nums2 []int
	out   int
}

var tests = []Test{
	{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, 5},
	{[]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}, 3},
	{[]int{1, 2, 4, 9, 8}, []int{3, 4, 2, 4, 8}, 2},
	{[]int{0, 1, 1, 1, 1}, []int{1, 0, 1, 0, 1}, 2},
}

func TestFindLength(t *testing.T) {
	for _, test := range tests {
		if res := findLength(test.nums1, test.nums2); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.nums1, test.nums2, test.out, res)
		}
	}
}

func BenchmarkFindLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := tests[i%4]
		findLength(t.nums1, t.nums2)
	}
}
