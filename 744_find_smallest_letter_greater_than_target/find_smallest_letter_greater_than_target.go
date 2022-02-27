package find_smallest_letter_greater_than_target

import (
	"sort"
)

func nextGreatestLetter(letters []byte, target byte) byte {
	pos := sort.Search(len(letters), func(i int) bool {
		return letters[i] >= target
	})

	for pos < len(letters) && letters[pos] == target {
		pos++
	}

	if pos == len(letters) {
		return letters[0]
	} else {
		return letters[pos]
	}
}
