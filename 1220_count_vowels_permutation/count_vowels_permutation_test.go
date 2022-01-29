package count_vowels_permutation

import "testing"

type Test struct {
	in  int
	out int
}

var tests = []Test{
	{1, 5},
	{2, 10},
	{3, 19},
	{4, 35},
	{5, 68},
	{8, 474},
	{32, 789289148},
	{76, 413022267},
	{1984, 504884281},
}

func TestCountVowelPermutation(t *testing.T) {
	for _, test := range tests {
		if res := countVowelPermutation(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}

func BenchmarkCountVowelPermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countVowelPermutation(32)
	}
}
