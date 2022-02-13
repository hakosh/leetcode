package meeting_rooms

import "testing"

type Test struct {
	intervals [][]int
	out       bool
}

var tests = []Test{
	{[][]int{{0, 30}, {5, 10}, {15, 20}}, false},
	{[][]int{{7, 10}, {2, 4}}, true},
	{[][]int{}, true},
}

func TestCanAttendMeetings(t *testing.T) {
	for _, test := range tests {
		if res := canAttendMeetings(test.intervals); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.intervals, test.out, res)
		}
	}
}
