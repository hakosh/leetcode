package minimum_deletions_to_make_character_frequencies_unique

import "testing"

type Test struct {
	s   string
	out int
}

var tests = []Test{
	{"a", 0},
	{"bbcebab", 2},
	{"ceabaacb", 2},
	{"aaabbbcc", 2},
	{"kenkvquhbxtzajcndjtujkryfzvjboucihsvnhdbngsamucxyiybjaecviyboamwotjfgjwvijvxvuffmwrvptsiqxzzejsfaageilyjjrjkvb", 68},
}

func TestMinDeletions(t *testing.T) {
	for _, test := range tests {
		if res := minDeletions(test.s); res != test.out {
			t.Errorf("For %q expected %d, got %d", test.s, test.out, res)
		}
	}
}

func BenchmarkMinDeletions(b *testing.B) {
	test := tests[len(tests)-1]

	for i := 0; i < b.N; i++ {
		minDeletions(test.s)
	}
}
