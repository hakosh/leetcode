package minimum_difficulty_of_a_job_schedule

import (
	"math"
)

var day int
var cache [][]int
var hardestFrom []int

func minDifficulty(jobDifficulty []int, d int) int {
	day = d

	if len(jobDifficulty) < d {
		return -1
	}

	hardestFrom = make([]int, len(jobDifficulty)+1)

	cache = make([][]int, len(jobDifficulty)+1)
	for i := range cache {
		cache[i] = make([]int, d+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	hardest := jobDifficulty[len(jobDifficulty)-1]
	for i := len(jobDifficulty) - 1; i >= 0; i-- {
		if jobDifficulty[i] > hardest {
			hardest = jobDifficulty[i]
		}

		hardestFrom[i] = hardest
	}

	return dp(jobDifficulty, 0, 1)
}

func dp(jobDifficulty []int, i, d int) int {
	if d == day {
		return hardestFrom[i]
	}

	if cache[i][d] == -1 {
		min := math.MaxInt32
		hardest := 0

		for j := i; j < len(jobDifficulty)-(day-d); j++ {
			if jobDifficulty[j] > hardest {
				hardest = jobDifficulty[j]
			}

			r := hardest + dp(jobDifficulty, j+1, d+1)

			if r < min {
				min = r
			}
		}

		cache[i][d] = min
	}

	return cache[i][d]
}
