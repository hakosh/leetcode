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
		grid := cloneGrid(test.grid)

		if res := numIslands(grid); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.grid, test.out, res)
		}
	}
}

func TestNumIslandsDFS(t *testing.T) {
	for _, test := range tests {
		grid := cloneGrid(test.grid)

		if res := numIslandsDFS(grid); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.grid, test.out, res)
		}
	}
}

func cloneGrid(grid [][]byte) [][]byte {
	rows := make([][]byte, len(grid))

	for i, row := range grid {
		rows[i] = append(rows[i], row...)
	}

	return rows
}

func BenchmarkNumIslands(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grid1 := cloneGrid(grid1)
		numIslands(grid1)

		grid2 := cloneGrid(grid2)
		numIslands(grid2)
	}
}

func BenchmarkNumIslandsDFS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grid1 := cloneGrid(grid1)
		numIslandsDFS(grid1)

		grid2 := cloneGrid(grid2)
		numIslandsDFS(grid2)
	}
}
