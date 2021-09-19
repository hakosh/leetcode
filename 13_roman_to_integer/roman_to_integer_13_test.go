package _3_roman_to_integer

import "testing"

type romanToIntTest struct {
	in string
	out int
}

var tests = []romanToIntTest{
	{ in: "III", out: 3 },
	{ in: "IV", out: 4 },
	{ in: "IX", out: 9 },
	{ in: "LVIII", out: 58 },
	{ in: "MCMXCIV", out: 1994 },
}

func TestRomanToInt(t *testing.T) {
	for _, test := range tests {
		result := romanToInt(test.in)

		if result != test.out {
			t.Errorf("For input %q expected %d, got %d", test.in, test.out, result)
		}
	}
}

func BenchmarkRomanToInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		romanToInt("MCMXCIV")
	}
}