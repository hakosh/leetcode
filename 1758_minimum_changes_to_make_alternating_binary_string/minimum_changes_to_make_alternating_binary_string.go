package minimum_changes_to_make_alternating_binary_string

func minOperations(s string) int {
	a, b := 0, 0
	s1, s2 := '0', '1'

	for _, c := range s {
		if c != s1 {
			a++
		}

		if c != s2 {
			b++
		}

		s1, s2 = s2, s1
	}

	if a < b {
		return a
	} else {
		return b
	}
}
