package maximum_score_from_performing_multiplication_operations

import "testing"

type Test struct {
	nums        []int
	multipliers []int
	out         int
}

var tests = []Test{
	{[]int{1, 2, 3}, []int{3, 2, 1}, 14},
	{[]int{-5, -3, -3, -2, 7, 1}, []int{-10, -5, 3, 4, 6}, 102},
}

func TestMaximumScore(t *testing.T) {
	for _, test := range tests {
		if res := maximumScore(test.nums, test.multipliers); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.nums, test.multipliers, test.out, res)
		}
	}
}

func BenchmarkMaximumScore(b *testing.B) {
	t := tests[1]

	for i := 0; i < b.N; i++ {
		maximumScore(t.nums, t.multipliers)
	}
}
