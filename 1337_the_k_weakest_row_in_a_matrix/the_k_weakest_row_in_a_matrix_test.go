package the_k_weakest_row_in_a_matrix

import (
	"fmt"
	"reflect"
	"testing"
)

type Test struct {
	mat  [][]int
	k    int
	want []int
}

var mat1 = [][]int{
	{1, 1, 0, 0, 0},
	{1, 1, 1, 1, 0},
	{1, 0, 0, 0, 0},
	{1, 1, 0, 0, 0},
	{1, 1, 1, 1, 1}}

var mat2 = [][]int{
	{1, 0, 0, 0},
	{1, 1, 1, 1},
	{1, 0, 0, 0},
	{1, 0, 0, 0}}

var mat3 = [][]int{{1, 1, 0}, {1, 1, 0}, {1, 1, 1}, {1, 1, 1}, {0, 0, 0}, {1, 1, 1}, {1, 0, 0}}

var tests = []Test{
	{mat1, 3, []int{2, 0, 3}},
	{mat2, 2, []int{0, 2}},
	{mat3, 6, []int{4, 6, 0, 1, 2, 3}},
}

func TestKWeakestRow(t *testing.T) {
	fmt.Println(strength([]int{1, 1, 1, 0, 0}))
	for _, test := range tests {
		if res := kWeakestRows(test.mat, test.k); !reflect.DeepEqual(res, test.want) {
			t.Errorf("for %v and %v wanted %v, got %v", test.mat, test.k, test.want, res)
		}
	}
}
