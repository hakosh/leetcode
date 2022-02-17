package diagonal_traverse

import (
	"reflect"
	"testing"
)

type Test struct {
	mat [][]int
	out []int
}

var square = [][]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
}

var rect1 = [][]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
	{10, 11, 12},
	{13, 14, 15},
}

var tests = []Test{
	{square, []int{1, 2, 4, 7, 5, 3, 6, 8, 9}},
	{rect1, []int{1, 2, 4, 7, 5, 3, 6, 8, 10, 13, 11, 9, 12, 14, 15}},
	{[][]int{{1}}, []int{1}},
}

func TestFindDiagonalOrder(t *testing.T) {
	for _, test := range tests {
		if res := findDiagonalOrder(test.mat); !reflect.DeepEqual(test.out, res) {
			t.Errorf("For %v expected %v, got %v", test.mat, test.out, res)
		}
	}
}
