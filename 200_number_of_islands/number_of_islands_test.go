package number_of_islands

import "testing"

type Test struct {
	grid [][]byte
	out  int
}

var grid1 = [][]byte{
	{'1', '1', '1', '1', '0'},
	{'1', '1', '0', '1', '0'},
	{'1', '1', '0', '0', '0'},
	{'0', '0', '0', '0', '0'},
}

var grid2 = [][]byte{
	{'1', '1', '0', '0', '0'},
	{'1', '1', '0', '0', '0'},
	{'0', '0', '1', '0', '0'},
	{'0', '0', '0', '1', '1'},
}

var tests = []Test{
	{grid1, 1},
	{grid2, 3},
}

func TestNumIslands(t *testing.T) {
	for _, test := range tests {
		if res := numIslands(test.grid); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.grid, test.out, res)
		}
	}
}
