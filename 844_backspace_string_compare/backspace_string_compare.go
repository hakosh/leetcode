package backspace_string_compare

func backspaceCompare(s string, t string) bool {
	p1, p2 := len(s)-1, len(t)-1

	for p1 >= 0 || p2 >= 0 {
		if p1 >= 0 && s[p1] == '#' {
			p1 = move(s, p1)
			continue
		}

		if p2 >= 0 && t[p2] == '#' {
			p2 = move(t, p2)
			continue
		}

		if (p1 < 0 && p2 >= 0) || (p1 >= 0 && p2 < 0) || s[p1] != t[p2] {
			return false
		}

		p1--
		p2--
	}

	return true
}

func move(str string, p int) int {
	q := p - 2
	p--

	for p > -1 && q < p {
		if str[p] == '#' {
			q -= 2
		}

		p--
	}

	return p
}
