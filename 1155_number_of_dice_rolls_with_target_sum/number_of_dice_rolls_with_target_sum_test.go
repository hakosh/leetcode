package number_of_dice_rolls_with_target_sum

import "testing"

type Test struct {
	n      int
	k      int
	target int
	out    int
}

var tests = []Test{
	{1, 6, 3, 1},
	{2, 6, 7, 6},
	{3, 6, 7, 15},
	{3, 16, 32, 150},
	{4, 16, 32, 2675},
	{30, 20, 193, 87756311},
	{30, 30, 500, 222616187},
}

func TestNumRollsToTarget(t *testing.T) {
	for _, test := range tests {
		if res := numRollsToTarget(test.n, test.k, test.target); res != test.out {
			t.Errorf("For %v, %v, %v expected %v, got %v", test.n, test.k, test.target, test.out, res)
		}
	}
}

func BenchmarkNumRollsToTarget(b *testing.B) {
	t := tests[5]

	for i := 0; i < b.N; i++ {
		numRollsToTarget(t.n, t.k, t.target)
	}
}
