package spiral_matrix_iii

import (
	"reflect"
	"testing"
)

type Test struct {
	rows   int
	cols   int
	rStart int
	cStart int
	want   [][]int
}

var tests = []Test{
	{1, 4, 0, 0, [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}}},
	{5, 6, 1, 4, [][]int{{1, 4}, {1, 5}, {2, 5}, {2, 4}, {2, 3}, {1, 3}, {0, 3}, {0, 4}, {0, 5}, {3, 5}, {3, 4}, {3, 3}, {3, 2}, {2, 2}, {1, 2}, {0, 2}, {4, 5}, {4, 4}, {4, 3}, {4, 2}, {4, 1}, {3, 1}, {2, 1}, {1, 1}, {0, 1}, {4, 0}, {3, 0}, {2, 0}, {1, 0}, {0, 0}}},
}

func TestSpiralMatrixIII(t *testing.T) {
	for i, test := range tests {
		if res := spiralMatrixIII(test.rows, test.cols, test.rStart, test.cStart); !reflect.DeepEqual(res, test.want) {
			t.Errorf("for test %d wanted %v, got %v", i, test.want, res)
		}
	}
}
