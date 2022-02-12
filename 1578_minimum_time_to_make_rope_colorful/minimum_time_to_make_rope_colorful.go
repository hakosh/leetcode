package minimum_time_to_make_rope_colorful

func minCost(colors string, neededTime []int) int {
	time := 0
	last := 0

	for i := 1; i < len(colors); i++ {
		if colors[last] != colors[i] {
			last = i
			continue
		}

		if neededTime[i] < neededTime[last] {
			time += neededTime[i]
		} else {
			time += neededTime[last]
			last = i
		}
	}

	return time
}
