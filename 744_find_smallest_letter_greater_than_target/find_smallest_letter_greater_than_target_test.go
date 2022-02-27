package find_smallest_letter_greater_than_target

import "testing"

type Test struct {
	letters []byte
	target  byte
	out     byte
}

var tests = []Test{
	{[]byte{'c', 'f', 'j'}, 'a', 'c'},
	{[]byte{'c', 'f', 'j'}, 'c', 'f'},
	{[]byte{'c', 'f', 'j'}, 'g', 'j'},
	{[]byte{'c', 'f', 'j'}, 'j', 'c'},
	{[]byte{'a', 'b'}, 'z', 'a'},
	{[]byte{'e', 'e', 'e', 'n', 'n'}, 'e', 'n'},
}

func TestNextGreatestLetter(t *testing.T) {
	for _, test := range tests {
		if res := nextGreatestLetter(test.letters, test.target); res != test.out {
			t.Errorf("For %q and %c expected %c, got %c", test.letters, test.target, test.out, res)
		}
	}
}
