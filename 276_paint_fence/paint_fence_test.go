package paint_fence

import "testing"

type Test struct {
	posts  int
	colors int
	out    int
}

var tests = []Test{
	{3, 2, 6},
	{7, 2, 42},
	{7, 8, 1950984},
	{12, 6, 1696518750},
	{2, 46340, 2147395600},
}

func TestNumWays(t *testing.T) {
	for _, test := range tests {
		if res := numWays(test.posts, test.colors); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.posts, test.colors, test.out, res)
		}
	}
}

func BenchmarkNumWays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numWays(12, 6)
	}
}
