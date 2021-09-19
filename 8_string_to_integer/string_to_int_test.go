package string_to_integer

import "testing"

type stringToIntTest struct {
	in  string
	out int
}

var tests = []stringToIntTest{
	{in: "42", out: 42},
	{in: "   -42", out: -42},
	{in: "+-12", out: 0},
	{in: "-+12", out: 0},
	{in: "   ", out: 0},
	{in: "    +0 123", out: 0},
	{in: "00000-42a1234", out: 0},
	{in: "4193 with words", out: 4193},
	{in: "words and 987", out: 0},
	{in: "-91283472332", out: -2147483648},
	{in: "20000000000000000000", out: 2147483647},
	{in: "  0000000000012345678", out: 12345678},
}

func TestStringToInt(t *testing.T) {
	for _, test := range tests {
		got := myAtoi(test.in)

		if got != test.out {
			t.Errorf("For %s expected %d, got %d\n", test.in, test.out, got)
		}
	}
}

func BenchmarkStringToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myAtoi("42")
	}
}

func BenchmarkStringToInt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myAtoi("4193 with words")
	}
}

func BenchmarkStringToInt3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myAtoi("  0000000000012345678")
	}
}

func BenchmarkStringToInt4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myAtoi("20000000000000000000")
	}
}

func BenchmarkStringToInt5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myAtoi("-91283472332")
	}
}