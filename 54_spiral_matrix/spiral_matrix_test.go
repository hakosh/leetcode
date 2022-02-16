package spiral_matrix

import (
	"reflect"
	"testing"
)

type Test struct {
	matrix [][]int
	out    []int
}

var matrix1 = [][]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
}

var matrix2 = [][]int{
	{1, 2, 3, 4},
	{5, 6, 7, 8},
	{9, 10, 11, 12},
	{13, 14, 15, 16},
}

var matrix3 = [][]int{
	{1, 2, 3, 4},
	{5, 6, 7, 8},
	{9, 10, 11, 12},
}

var matrix4 = [][]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
	{10, 11, 12},
}

var tests = []Test{
	{matrix1, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
	{matrix2, []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10}},
	{matrix3, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	{matrix4, []int{1, 2, 3, 6, 9, 12, 11, 10, 7, 4, 5, 8}},
}

func TestSpiralOrder(t *testing.T) {
	for _, test := range tests {
		if res := spiralOrder(test.matrix); !reflect.DeepEqual(res, test.out) {
			t.Errorf("For %v expected %v, got %v", test.matrix, test.out, res)
		}
	}
}
