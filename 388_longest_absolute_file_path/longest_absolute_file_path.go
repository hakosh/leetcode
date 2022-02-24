package longest_absolute_file_path

import "strings"

func lengthLongestPath(input string) int {
	path := make([]string, 0)
	lines := strings.Split(input, "\n")

	longest := 0

	for _, line := range lines {
		d := depth(line)
		path = append(path[0:d], line[d:])

		if this := length(path); this > longest {
			longest = this
		}
	}

	return longest
}

func length(path []string) int {
	count := 0
	isFile := false

	for _, part := range path {
		for _, c := range part {
			if c == '.' {
				isFile = true
			}

			count++
		}

		count++
	}

	if isFile {
		return count - 1
	} else {
		return 0
	}
}

func depth(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] != '\t' {
			return i
		}
	}

	return 0
}
