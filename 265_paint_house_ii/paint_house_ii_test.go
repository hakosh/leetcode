package paint_house_ii

import "testing"

type Test struct {
	in  [][]int
	out int
}

var tests = []Test{

	{[][]int{{1, 3}, {2, 4}}, 5},
	{[][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 9}}, 10},
	{[][]int{{7, 6, 2}}, 2},
	{[][]int{{3, 5, 3}, {6, 17, 6}, {7, 13, 18}, {9, 10, 19}}, 26},
	{[][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}, {16, 9, 18}, {9, 5, 13}, {8, 13, 20}, {19, 14, 7}}, 46},
}

func TestMinCostII(t *testing.T) {
	for _, test := range tests {
		if res := minCostII(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}

func BenchmarkMinCostII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minCostII(tests[3].in)
	}
}
