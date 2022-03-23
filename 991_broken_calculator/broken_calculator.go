package broken_calculator

func brokenCalc(startValue int, target int) int {
	steps := 0

	for startValue < target {
		steps++

		if target%2 == 1 {
			target++
		} else {
			target /= 2
		}
	}

	return steps + startValue - target
}
