package remove_duplicate_letters

import "testing"

type Test struct {
	s    string
	want string
}

var tests = []Test{
	{"bcabc", "abc"},
	{"cbacdcbc", "acdb"},
	{"abacb", "abc"},
	{"fahkljafhkfdyilujahbwrem", "afhkdyilujbwrem"},
}

func TestRemoveDuplicateLetters(t *testing.T) {
	for _, test := range tests {
		if res := removeDuplicateLetters(test.s); res != test.want {
			t.Errorf("for %q wanted %q, got %q", test.s, test.want, res)
		}
	}
}
