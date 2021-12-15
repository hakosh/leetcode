package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	letters := make([]int, 256)

	for _, char := range magazine {
		letters[char]++
	}

	for _, char := range ransomNote {
		if letters[char] <= 0 {
			return false
		}

		letters[char]--
	}

	return true
}
