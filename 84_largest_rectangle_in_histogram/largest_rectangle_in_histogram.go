package largest_rectangle_in_histogram

func largestRectangleArea(heights []int) int {
	area := 0
	stack := make([]int, 1)
	stack[0] = -1

	for i := 0; i < len(heights); i++ {
		for len(stack) > 1 && heights[stack[len(stack)-1]] >= heights[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			this := heights[top] * (i - stack[len(stack)-1] - 1)
			if this > area {
				area = this
			}
		}

		stack = append(stack, i)
	}

	for len(stack) > 1 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		this := heights[top] * (len(heights) - stack[len(stack)-1] - 1)
		if this > area {
			area = this
		}
	}

	return area
}

func largestRectangleArea2(heights []int) int {
	area := 0
	stack := make([][2]int, 1, 16)
	stack[0] = [2]int{0, -1}

	for i := 0; i <= len(heights); i++ {
		height := 0
		if i < len(heights) {
			height = heights[i]
		}

		top := stack[len(stack)-1]

		if height > top[1] {
			stack = append(stack, [2]int{i + 1, height})
		} else {
			for top[1] >= height {
				this := (i - top[0] + 1) * top[1]

				if this > area {
					area = this
				}

				top = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}

			stack = append(stack, top, [2]int{top[0] + 1, height}, [2]int{i + 1, height})
		}
	}

	return area
}
