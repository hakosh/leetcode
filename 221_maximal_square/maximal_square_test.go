package maximal_square

import "testing"

type Test struct {
	matrix [][]byte
	out    int
}

var matrix = [][]byte{
	{'1', '0', '1', '0', '0'},
	{'1', '0', '1', '1', '1'},
	{'1', '1', '1', '1', '1'},
	{'1', '0', '0', '1', '0'},
}

var matrix2 = [][]byte{
	{'1', '1', '1', '0', '0'},
	{'1', '1', '1', '0', '0'},
	{'1', '1', '1', '0', '0'},
	{'0', '0', '0', '1', '1'},
	{'0', '0', '0', '1', '1'},
	{'0', '0', '0', '0', '1'},
	{'0', '0', '0', '0', '1'},
}

var matrix3 = [][]byte{
	{'1', '1', '0', '1'},
	{'1', '1', '0', '1'},
	{'1', '1', '1', '1'},
}

var tests = []Test{
	{[][]byte{{'0', '1'}}, 1},
	{[][]byte{{'0'}, {'1'}, {'1'}}, 1},
	{[][]byte{{'1', '0'}, {'0', '1'}}, 1},
	{matrix, 4},
	{matrix2, 9},
	{matrix3, 4},
}

func TestMaximalSquare(t *testing.T) {
	for _, test := range tests {
		res := maximalSquare(test.matrix)

		if res != test.out {
			t.Errorf("For %v wanted %v, got %v\n", test.matrix, test.out, res)
		}
	}
}

func BenchmarkMaximalSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maximalSquare(matrix2)
	}
}
