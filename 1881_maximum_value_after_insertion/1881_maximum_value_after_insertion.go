package maximum_value_after_insertion

func maxValue(n string, x int) string {
	digit := rune(x + '0')

	if n[0] == '-' {
		for i, c := range n {
			if c > digit {
				return n[:i] + string(digit) + n[i:]
			}
		}
	} else {
		for i, c := range n {
			if c < digit {
				return n[:i] + string(digit) + n[i:]
			}
		}
	}

	return n + string(digit)
}
