package keys_and_rooms

import "testing"

type Test struct {
	in  [][]int
	out bool
}

var tests = []Test{
	{[][]int{{1}, {2}, {3}, {}}, true},
	{[][]int{{1, 3}, {3, 0, 1}, {2}, {0}}, false},
}

func TestCanVisitAllRooms(t *testing.T) {
	for _, test := range tests {
		if res := canVisitAllRooms(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
