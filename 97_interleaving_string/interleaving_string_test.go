package interleaving_string

import "testing"

type Test struct {
	s1  string
	s2  string
	s3  string
	out bool
}

var tests = []Test{
	{"a", "b", "a", false},
	{"a", "", "c", false},
	{"aabc", "abad", "aabcabad", true},
	{"aabccc", "dbbcaba", "aadbbcbcaccba", true},
	{"aabccxy", "dbbcaui", "aadbbbacccxyui", false},
}

func TestIsInterleve(t *testing.T) {
	for _, test := range tests {
		if res := isInterleave(test.s1, test.s2, test.s3); res != test.out {
			t.Errorf("For %v, %v and %v expected %v, got %v", test.s1, test.s2, test.s3, test.out, res)
		}
	}
}

func BenchmarkIsInterleave(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := tests[i%len(tests)]
		isInterleave(t.s1, t.s2, t.s3)
	}
}
