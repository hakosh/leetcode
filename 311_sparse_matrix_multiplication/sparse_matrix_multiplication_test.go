package sparse_matrix_multiplication

import (
	"reflect"
	"testing"
)

type Test struct {
	mat1 [][]int
	mat2 [][]int
	want [][]int
}

var tests = []Test{
	{[][]int{{1, 0, 0}, {-1, 0, 3}}, [][]int{{7, 0, 0}, {0, 0, 0}, {0, 0, 1}}, [][]int{{7, 0, 0}, {-7, 0, 3}}},
	{[][]int{{0}}, [][]int{{0}}, [][]int{{0}}},
}

func TestMultiply(t *testing.T) {
	for _, test := range tests {
		if res := multiply(test.mat1, test.mat2); !reflect.DeepEqual(res, test.want) {
			t.Errorf("for %v and %v wanted %v, got %v", test.mat1, test.mat2, test.want, res)
		}
	}
}
