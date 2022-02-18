package merge_intervals

import (
	"reflect"
	"testing"
)

type Test struct {
	intervals [][]int
	out       [][]int
}

var tests = []Test{
	{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
	{[][]int{{1, 3}, {2, 6}, {9, 10}}, [][]int{{1, 6}, {9, 10}}},
	{[][]int{{1, 3}}, [][]int{{1, 3}}},
	{[][]int{{1, 4}, {2, 3}}, [][]int{{1, 4}}},
}

func TestMerge(t *testing.T) {
	for _, test := range tests {
		if res := merge(test.intervals); !reflect.DeepEqual(res, test.out) {
			t.Errorf("For %v expected %v, got %v", test.intervals, test.out, res)
		}
	}
}
