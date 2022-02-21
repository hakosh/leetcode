package largest_rectangle_in_histogram

import "testing"

type Test struct {
	heights []int
	out     int
}

var tests = []Test{
	{[]int{2, 1, 5, 6, 2, 3}, 10},
	{[]int{2, 1, 4, 5, 4, 2, 2, 3, 2, 2}, 16},
	{[]int{2, 4}, 4},
	{[]int{1, 1}, 2},
	{[]int{2, 1, 2}, 3},
	{[]int{0, 0, 2, 1, 2}, 3},
}

func TestLargestRectangleArea(t *testing.T) {
	for _, test := range tests {
		if res := largestRectangleArea(test.heights); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.heights, test.out, res)
		}
	}
}

func BenchmarkLargestRectangleArea(b *testing.B) {
	in := []int{2, 1, 4, 5, 4, 2, 2, 3, 2, 2, 5, 0, 5, 0, 3}

	for i := 0; i < b.N; i++ {
		largestRectangleArea(in)
	}
}
