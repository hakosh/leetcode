package reverse_words_in_string_ii

import (
	"reflect"
	"testing"
)

type Test struct {
	s   []byte
	out []byte
}

var tests = []Test{
	{[]byte{'t', 'h', 'e', ' ', 's', 'k', 'y', ' ', 'i', 's', ' ', 'b', 'l', 'u', 'e'}, []byte{'b', 'l', 'u', 'e', ' ', 'i', 's', ' ', 's', 'k', 'y', ' ', 't', 'h', 'e'}},
	{[]byte{'m', 'a', 'j', 'o', 'm'}, []byte{'m', 'a', 'j', 'o', 'm'}},
}

func TestReverseWords(t *testing.T) {
	for _, test := range tests {
		res := make([]byte, len(test.s))
		copy(res, test.s)
		reverseWords(res)

		if !reflect.DeepEqual(res, test.out) {
			t.Errorf("For %q expected %q, got %q", string(test.s), string(test.out), string(res))
		}
	}
}
