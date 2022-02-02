package implement_strstr

import "testing"

type Test struct {
	haystack string
	needle   string
	out      int
}

var tests = []Test{
	{"hello", "ll", 2},
	{"hi", "", 0},
	{"aaa", "aaaa", -1},
	{"monkey", "money", -1},
	{"", "", 0},
	{"x", "x", 0},
	{"xbc", "x", 0},
	{"xbc", "c", 2},
	{"mississippi", "issip", 4},
}

func TestStrStr(t *testing.T) {
	for _, test := range tests {
		if res := strStr(test.haystack, test.needle); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.haystack, test.needle, test.out, res)
		}
	}
}
