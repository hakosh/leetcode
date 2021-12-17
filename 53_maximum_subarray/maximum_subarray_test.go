package maximum_subarray

import (
	"testing"
)

type Test struct {
	in   []int
	want int
}

var cases = []Test{
	{in: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, want: 6},
	{in: []int{1}, want: 1},
	{in: []int{-2, 1}, want: 1},
	{in: []int{5, 4, -1, 7, 8}, want: 23},
}

func TestThirdMax(t *testing.T) {
	for _, test := range cases {
		if res := maxSubArray(test.in); res != test.want {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.want, res)
		}
	}
}

func BenchmarkThirdMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	}
}
