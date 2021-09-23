package integer_to_roman

import "testing"

type Test struct {
	in int
	want string
}

var cases = []Test{
	{ in: 3, want: "III" },
	{ in: 4, want: "IV" },
	{ in: 9, want: "IX" },
	{ in: 58, want: "LVIII" },
	{ in: 1994, want: "MCMXCIV"},
}

func TestMoveZeroes(t *testing.T) {
	for _, test := range cases {
		result := integerToRoman(test.in)

		if result != test.want {
			t.Errorf("For %v wanted %v, got %v\n", test.in, test.want, result)
		}
	}
}
