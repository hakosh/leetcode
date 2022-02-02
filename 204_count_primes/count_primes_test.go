package count_primes

import "testing"

type Test struct {
	in  int
	out int
}

var tests = []Test{
	{10, 4},
	{1000, 168},
	{5000, 669},
	{0, 0},
	{1, 0},
}

func TestCountPrimes(t *testing.T) {
	for _, test := range tests {
		if res := countPrimes(test.in); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.in, test.out, res)
		}
	}
}

func BenchmarkCountPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countPrimes(5000000)
	}
}
