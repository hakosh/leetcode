package reverse_words_in_string

import "testing"

type Test struct {
	s   string
	out string
}

var tests = []Test{
	{"the sky is blue", "blue is sky the"},
	{"the sky   is    blue", "blue is sky the"},
	{"majom", "majom"},
	{"   hello world   ", "world hello"},
	{"a good    example", "example good a"},
}

func TestReverseWords(t *testing.T) {
	for _, test := range tests {
		if res := reverseWords(test.s); res != test.out {
			t.Errorf("For %q expected %q, got %q", test.s, test.out, res)
		}
	}
}
