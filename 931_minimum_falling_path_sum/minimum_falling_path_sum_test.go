package minimum_falling_path_sum

import "testing"

type Test struct {
	grid [][]int
	out  int
}

var grid1 = [][]int{
	{2, 1, 3},
	{6, 5, 4},
	{7, 8, 9},
}

var grid2 = [][]int{
	{-19, 57},
	{-40, -5},
}

var tests = []Test{
	{grid1, 13},
	{grid2, -59},
}

func TestMinFallingPathSum(t *testing.T) {
	for _, test := range tests {
		if res := minFallingPathSum(test.grid); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.grid, test.out, res)
		}
	}
}
