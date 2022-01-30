package domino_and_tromino_tiling

import "testing"

type Test struct {
	n   int
	out int
}

var tests = []Test{
	{3, 5},
	{1, 1},
	{5, 24},
	{10, 1255},
	{15, 65501},
	{75, 231534077},
	{173, 437507194},
}

func TestNumTilings(t *testing.T) {
	for _, test := range tests {
		if res := numTilings(test.n); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.n, test.out, res)
		}
	}
}

func BenchmarkNumTilings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numTilings(75)
	}
}
