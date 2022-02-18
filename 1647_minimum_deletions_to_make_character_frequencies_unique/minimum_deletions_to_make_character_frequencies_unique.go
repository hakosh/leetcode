package minimum_deletions_to_make_character_frequencies_unique

func minDeletions(s string) int {
	max := 0
	steps := 0

	// count chars
	count := [27]int{}
	for i := range s {
		char := s[i] - 'a'
		count[char]++

		if count[char] > max {
			max = count[char]
		}
	}

	// group by count
	groups := make([]bool, max+1)

	for _, count := range count {
		if count == 0 {
			continue
		}

		check := count
		for check > 0 && groups[check] {
			check--
		}

		groups[check] = true
		steps += count - check
	}

	return steps
}
