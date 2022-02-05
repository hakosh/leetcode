package longest_substring_without_repeating_characters

import "testing"

type Test struct {
	in  string
	out int
}

var tests = []Test{
	{"", 0},
	{"aab", 2},
	{"dvdf", 3},
	{"bbbbbbb", 1},
	{"abbbc", 2},
	{"abxbc", 3},
	{"pwwkew", 3},
	{"abcabcbb", 3},
	{"majomm", 5},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, test := range tests {
		if res := lengthOfLongestSubstring(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
