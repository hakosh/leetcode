package daily_temperatures

func dailyTemperatures(temperatures []int) []int {
	days := make([]int, len(temperatures))
	stack := make([]int, 0, 8)

	for i, this := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < this {
			pos := stack[len(stack)-1]
			days[pos] = i - pos
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return days
}
