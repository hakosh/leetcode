package reverse_bits

import "testing"

type Test struct {
	in  uint32
	out uint32
}

var tests = []Test{
	{43261596, 964176192},
	{4294967293, 3221225471},
	{0, 0},
	{1, 2147483648},
}

func TestReverseBits(t *testing.T) {
	for _, test := range tests {
		if res := reverseBits(test.in); res != test.out {
			t.Errorf("For %b expected %b, got %b", test.in, test.out, res)
		}
	}
}
