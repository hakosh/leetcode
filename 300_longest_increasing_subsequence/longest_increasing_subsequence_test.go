package longest_increasing_subsequence

import "testing"

type Test struct {
	in  []int
	out int
}

var tests = []Test{
	{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
	{[]int{10, 9, 2, 5, 3, 7, 101, 18, 10, 20, 22}, 6},
	{[]int{0, 1, 0, 3, 2, 3}, 4},
	{[]int{3}, 1},
	{[]int{1, 1, 1}, 1},
	{[]int{-2, -1}, 2},
	{[]int{1, 3, 6, 7, 9, 4, 10, 5, 6}, 6},
}

func TestLengthOfLIS(t *testing.T) {
	for _, test := range tests {
		if res := lengthOfLIS(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}

func BenchmarkLengthOfLIS(b *testing.B) {
	test := tests[1]

	for i := 0; i < b.N; i++ {
		lengthOfLIS(test.in)
	}
}
