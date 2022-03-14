package strobogrammatic_number_ii

var wraps = []string{"00", "11", "69", "88", "96"}

func findStrobogrammatic(n int) []string {
	return gen(n, true)
}

func gen(n int, top bool) []string {
	if n == 0 {
		return []string{}
	}

	if n == 1 {
		return []string{"0", "1", "8"}
	}

	inner := gen(n-2, false)
	list := make([]string, 0, len(inner)+5)

	for _, sides := range wraps {
		if top && sides[0] == '0' {
			continue
		}

		if len(inner) == 0 {
			list = append(list, sides)
			continue
		}

		for _, str := range inner {
			list = append(list, sides[:1]+str+sides[1:])
		}
	}

	return list
}
