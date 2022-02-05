package trapping_rain_waiter

import "testing"

type Test struct {
	in  []int
	out int
}

var tests = []Test{
	{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
	{[]int{4, 2, 0, 3, 2, 5}, 9},
	{[]int{1, 2, 0, 3}, 2},
	{[]int{2, 1, 0, 3}, 3},
	{[]int{1}, 0},
	{[]int{0, 0, 0, 0}, 0},
	{[]int{3, 4, 5}, 0},
	{[]int{0, 50, 2, 1, 15, 3, 2, 1, 4, 7, 4, 3, 0, 0, 0, 8, 4, 3}, 83},
}

func TestTrap(t *testing.T) {
	for _, test := range tests {
		if res := trap(test.in); res != test.out {
			t.Errorf("For %v expected %d, got %d", test.in, test.out, res)
		}
	}
}

func BenchmarkTrap(b *testing.B) {
	test := tests[len(tests)-1]

	for i := 0; i < b.N; i++ {
		trap(test.in)
	}
}
