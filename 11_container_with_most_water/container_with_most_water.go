package container_with_most_water

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	area := 0

	for l != r {
		this := (r - l) * min(height[l], height[r])

		if this > area {
			area = this
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return area
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
