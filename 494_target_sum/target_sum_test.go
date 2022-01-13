package target_sum

import "testing"

type Test struct {
	in     []int
	target int
	out    int
}

var tests = []Test{
	{[]int{1, 1, 1, 1, 1}, 3, 5},
}

func TestFindTargetSumWays(t *testing.T) {
	for _, test := range tests {
		if res := findTargetSumWays(test.in, test.target); res != test.out {
			t.Errorf("For %v and %v exptected %v, got %v", test.in, test.target, test.out, res)
		}
	}
}

func BenchmarkFindTargetSumWays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findTargetSumWays(tests[0].in, 3)
	}
}
