package open_the_lock

import "testing"

type Test struct {
	deadends []string
	target   string
	out      int
}

var tests = []Test{
	{[]string{"0201", "0101", "0102", "1212", "2002"}, "0202", 6},
	{[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888", -1},
	{[]string{"8888"}, "0009", 1},
	{[]string{"0000"}, "8888", -1},
	{[]string{"9999"}, "0000", 0},
}

func TestOpenLock(t *testing.T) {
	for _, test := range tests {
		if res := openLock(test.deadends, test.target); res != test.out {
			t.Errorf("For %v and %v wanted %v, got %v", test.deadends, test.target, test.out, res)
		}
	}
}
