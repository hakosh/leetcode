package online_election

import "testing"

var queries = [][2]int{
	{3, 0},
	{12, 1},
	{25, 1},
	{15, 0},
	{24, 0},
	{8, 1},
}

func TestTopVotedCandidate(t *testing.T) {
	votes := Constructor([]int{0, 1, 1, 0, 0, 1, 0}, []int{0, 5, 10, 15, 20, 25, 30})

	for _, time := range queries {
		if res := votes.Q(time[0]); res != time[1] {
			t.Errorf("At %v wanted %v, got %v", time[0], time[1], res)
		}
	}
}

func BenchmarkTopVotedCandidate_Q(b *testing.B) {
	for i := 0; i < b.N; i++ {
		votes := Constructor([]int{0, 1, 1, 0, 0, 1, 0}, []int{0, 5, 10, 15, 20, 25, 30})

		for _, time := range queries {
			votes.Q(time[0])
		}
	}
}