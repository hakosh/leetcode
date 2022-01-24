package minimum_difficulty_of_a_job_schedule

import "testing"

type Test struct {
	difficulty []int
	days       int
	out        int
}

var tests = []Test{
	{[]int{6, 5, 4, 3, 2, 1}, 2, 7},
	{[]int{6, 5, 4, 3, 2, 1}, 7, -1},
	{[]int{1, 1, 1}, 3, 3},
	{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 10, 0},
	{[]int{11, 111, 22, 222, 33, 333, 44, 444}, 6, 843},
}

func TestMinDifficulty(t *testing.T) {
	for _, test := range tests {
		if res := minDifficulty(test.difficulty, test.days); res != test.out {
			t.Errorf("For %v and %v expected %v, got %v", test.difficulty, test.days, test.out, res)
		}
	}
}

func BenchmarkMinDifficulty(b *testing.B) {
	t := tests[0]

	for i := 0; i < b.N; i++ {
		minDifficulty(t.difficulty, t.days)
	}
}
