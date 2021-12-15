package first_unique_char

func firstUniqChar(s string) int {
	count := make([]int, 127)
	first := make([]int, 127)

	for i, char := range s {
		count[char]++

		if first[char] == 0 {
			first[char] = i
		}
	}

	firstUnique := len(s)

	for i := 97; i <= 122; i++ {
		if count[i] == 1 && first[i] < firstUnique {
			firstUnique = first[i]
		}
	}

	if firstUnique == len(s) {
		return -1
	}

	return firstUnique
}
