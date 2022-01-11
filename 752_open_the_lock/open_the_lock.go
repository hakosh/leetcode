package open_the_lock

func openLock(deadends []string, target string) int {
	queue := []string{"0000"}
	steps := 0

	stop := make(map[string]bool)
	for _, deadend := range deadends {
		stop[deadend] = true
	}

	for len(queue) > 0 {
		var tmp []string

		for _, cmb := range queue {
			if _, ok := stop[cmb]; ok {
				continue
			}

			if cmb == target {
				return steps
			}

			stop[cmb] = true

			for i := 0; i < 4; i++ {
				c := cmb[i]

				up := cmb[:i] + flipUp(c) + cmb[i+1:]
				down := cmb[:i] + flipDown(c) + cmb[i+1:]

				tmp = append(tmp, up, down)
			}
		}

		steps++
		queue = tmp
	}

	return -1
}

func flipUp(c byte) string {
	if c == '9' {
		return "0"
	}

	return string(c + 1)
}

func flipDown(c byte) string {
	if c == '0' {
		return "9"
	}

	return string(c - 1)
}
