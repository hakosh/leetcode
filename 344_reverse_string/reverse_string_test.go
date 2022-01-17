package reverse_string

import (
	"reflect"
	"testing"
)

type Test struct {
	in  []byte
	out []byte
}

var tests = []Test{
	{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
}

func TestDecodeString(t *testing.T) {
	for _, test := range tests {
		if reverseString(test.in); !reflect.DeepEqual(test.in, test.out) {
			t.Errorf("Expected %v, got %v", test.out, test.in)
		}
	}
}
