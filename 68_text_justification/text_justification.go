package text_justification

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	lines := groupWords(words, maxWidth)

	justified := make([]string, len(lines))
	for i, line := range lines {
		justified[i] = justify(line, maxWidth, i == len(lines)-1)
	}

	return justified
}

func groupWords(words []string, width int) [][]string {
	lines := make([][]string, 0)

	idx := -1
	cur := width

	for _, word := range words {
		if cur+len(word) > width {
			idx++
			cur = 0
			lines = append(lines, []string{})
		}

		lines[idx] = append(lines[idx], word)
		cur += len(word) + 1
	}

	return lines
}

func justify(line []string, width int, last bool) string {
	if last || len(line) == 1 {
		str := strings.Join(line, " ")
		return str + fmt.Sprintf("%*s", width-len(str), "")
	}

	chars := 0
	for _, word := range line {
		chars += len(word)
	}

	numWords := len(line) - 1
	gaps := make([][]byte, numWords)
	for i := 0; i < width-chars; i++ {
		gaps[i%len(gaps)] = append(gaps[i%len(gaps)], ' ')
	}

	s := strings.Builder{}
	for i, word := range line {
		s.WriteString(word)

		if i < len(line)-1 {
			s.Write(gaps[i])
		}
	}

	return s.String()
}
