package minimum_number_of_steps_to_make_two_strings_anagram

import "testing"

type Test struct {
	s   string
	t   string
	out int
}

var tests = []Test{
	{"leetcode", "practice", 5},
}

func TestMinSteps(t *testing.T) {
	for _, test := range tests {
		if res := minSteps(test.s, test.t); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.s, test.t, test.out, res)
		}
	}
}
