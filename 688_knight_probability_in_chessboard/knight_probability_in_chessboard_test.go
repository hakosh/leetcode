package knight_probability_in_chessboard

import "testing"

type Test struct {
	n   int
	k   int
	row int
	col int
	out float64
}

var tests = []Test{
	{3, 2, 0, 0, 0.0625},
	{1, 0, 0, 0, 1.0},
	{8, 2, 3, 3, 0.875},
	{8, 4, 3, 3, 0.48291015625},
}

func TestKnightProbability(t *testing.T) {
	for i, test := range tests {
		if res := knightProbability(test.n, test.k, test.row, test.col); res != test.out {
			t.Errorf("For case %d expected %v, got %v", i, test.out, res)
		}
	}
}

func BenchmarkKnightProbability(b *testing.B) {
	test := tests[len(tests)-1]

	for i := 0; i < b.N; i++ {
		knightProbability(test.n, test.k, test.row, test.col)
	}
}
