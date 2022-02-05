package trapping_rain_waiter

func trap(height []int) int {
	return trapFast(height)
}

func trapFast(height []int) int {
	total := 0

	rightMax := make([]int, len(height))
	rightMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i+1])
	}

	for i := 0; i < len(height)-1; {
		current := 0
		closedAt := 0

		left := height[i]
		right := rightMax[i]

		for j := i + 1; j < len(height); j++ {
			if height[j] >= left || height[j] == right {
				total += current
				closedAt = j
				break
			} else {
				current += min(left, right) - height[j]
			}
		}

		i = closedAt
	}

	return total
}

func trapDP(height []int) int {
	total := 0
	size := len(height)

	rightMax := make([]int, size)
	rightMax[size-1] = height[size-1]
	for i := size - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	leftMax := make([]int, size)
	leftMax[0] = height[0]
	for i := 1; i < size; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	for i := 1; i < size-1; i++ {
		total += min(leftMax[i], rightMax[i]) - height[i]
	}

	return total
}

func trapBruteForce(height []int) int {
	volume := 0

	for i, left := range height {
		for k := 0; k < left; k++ {
			distance := 0
			l := left - k

			for _, right := range height[i+1:] {
				r := right

				if r >= l {
					volume += distance
					break
				} else {
					distance++
				}
			}
		}
	}

	return volume
}

func trapStack(height []int) int {
	stack := make([]int, 0)
	volume := 0

	for i := range height {
		for len(stack) != 0 && height[i] > height[stack[len(stack)-1]] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				break
			}

			top := stack[len(stack)-1]

			distance := i - top - 1
			height := min(height[i], height[top]) - height[pop]

			volume += distance * height
		}

		stack = append(stack, i)
	}

	return volume
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
