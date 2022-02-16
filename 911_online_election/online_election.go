package online_election

import "sort"

type TopVotedCandidate struct {
	timeline [][2]int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	votes := make(map[int]int)
	timeline := make([][2]int, len(times))

	leadingCount := 0
	leader := persons[0]

	for i := 0; i < len(persons); i++ {
		person := persons[i]
		time := times[i]

		votes[person]++
		numVotes, _ := votes[person]

		if numVotes >= leadingCount {
			leadingCount = numVotes
			leader = person
		}

		timeline[i] = [2]int{time, leader}
	}

	return TopVotedCandidate{
		timeline: timeline,
	}
}

func (this *TopVotedCandidate) Q(t int) int {
	pos := sort.Search(len(this.timeline), func(i int) bool {
		return this.timeline[i][0] >= t
	})

	if pos < len(this.timeline) && (pos == 0 || this.timeline[pos][0] == t) {
		return this.timeline[pos][1]
	} else {
		return this.timeline[pos-1][1]
	}
}
