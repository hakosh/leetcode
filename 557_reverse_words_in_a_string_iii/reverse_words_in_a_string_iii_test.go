package reverse_words_in_a_string_iii

import "testing"

type Test struct {
	s   string
	out string
}

var tests = []Test{
	{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
}

func TestReverseWords(t *testing.T) {
	for _, test := range tests {
		if res := reverseWords(test.s); res != test.out {
			t.Errorf("For %q expected %q, got %q", test.s, test.out, res)
		}
	}
}
