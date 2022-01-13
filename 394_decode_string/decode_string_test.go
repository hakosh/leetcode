package decode_numbers

import "testing"

type Test struct {
	in  string
	out string
}

var tests = []Test{
	{"3[a]2[bc]", "aaabcbc"},
	{"3[a2[c]]", "accaccacc"},
	{"2[abc]3[cd]ef", "abcabccdcdcdef"},
	{"majom", "majom"},
	{"[x2[o]]", "xoo"},
}

func TestDecodeString(t *testing.T) {
	for _, test := range tests {
		if res := decodeString(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}
