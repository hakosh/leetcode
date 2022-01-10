package walls_and_gates

import (
	"math"
	"reflect"
	"testing"
)

type Test struct {
	room [][]int
	out  [][]int
}

var room1 = [][]int{
	{math.MaxInt32, -1, 0, math.MaxInt32},
	{math.MaxInt32, math.MaxInt32, math.MaxInt32, -1},
	{math.MaxInt32, -1, math.MaxInt32, -1},
	{0, -1, math.MaxInt32, math.MaxInt32},
}

var out1 = [][]int{
	{3, -1, 0, 1},
	{2, 2, 1, -1},
	{1, -1, 2, -1},
	{0, -1, 3, 4},
}

var tests = []Test{
	{room1, out1},
	{[][]int{{-1}}, [][]int{{-1}}},
}

func TestWallsAndGates(t *testing.T) {
	for _, test := range tests {
		var room [][]int

		for _, row := range test.room {
			newRow := make([]int, len(row))
			copy(newRow, row)
			room = append(room, newRow)
		}

		wallsAndGates(room)

		if !reflect.DeepEqual(room, test.out) {
			t.Errorf("Expected %v, got %v", test.out, room)
		}
	}
}
