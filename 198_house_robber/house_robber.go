package house_robber

var mem map[int]int

func rob(nums []int) int {
	mem = make(map[int]int, len(nums))
	return robInner(nums)
}

func robInner(nums []int) int {
	i := len(nums)

	if i == 1 {
		return nums[0]
	}

	if i == 2 {
		return max(nums[0], nums[1])
	}

	if v, ok := mem[i]; ok {
		return v
	} else {
		mem[i] = max(
			rob(nums[:i-1]),
			rob(nums[:i-2])+nums[i-1],
		)

		return mem[i]
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
