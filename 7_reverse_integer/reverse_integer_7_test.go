package __reverse_integer

import "testing"

type reverseTest struct {
	in, out int
}

var reverseTests = []reverseTest{
	{0, 0},
	{120, 21},
	{321, 123},
	{-407, -704},
	{1534236469, 0},
}

func TestReverse(t *testing.T) {
	for _, test := range reverseTests {
		if res := reverse(test.in); test.out != res {
			t.Errorf("Expected %d. got %d\n", test.out, res)
		}
	}
}
