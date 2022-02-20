package partition_labels

import "strings"

func partitionLabels(s string) []int {
	borders := make([]int, 0)

	for len(s) > 0 {
		pos := strings.LastIndexByte(s, s[0]) + 1

		for i := 1; i < pos; i++ {
			this := strings.LastIndexByte(s, s[i]) + 1

			if this > pos {
				pos = this
			}
		}

		borders = append(borders, pos)
		s = s[pos:]
	}

	return borders
}
