package is_anagram

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letters := make(map[rune]int)

	for _, char := range s {
		letters[char]++
	}

	for _, char := range t {
		count, ok := letters[char]

		if !ok || count <= 0 {
			return false
		}

		letters[char]--
	}

	return true
}
