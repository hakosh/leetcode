package unique_paths

import "testing"

type Test struct {
	m   int
	n   int
	out int
}

var tests = []Test{
	{3, 7, 28},
	//{3, 2, 3},
	//{9, 18, 1081575},
}

func TestUniquePaths(t *testing.T) {
	for _, test := range tests {
		if res := uniquePaths(test.m, test.n); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.m, test.n, test.out, res)
		}
	}
}

func BenchmarkUniquePaths(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uniquePaths(9, 18)
	}
}
