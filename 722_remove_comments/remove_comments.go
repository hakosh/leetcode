package remove_comments

import (
	"strings"
)

func removeComments(source []string) []string {
	clean := make([]string, 0)
	blockOpen := false

	line := strings.Builder{}

	for i := 0; i < len(source); i++ {
		for j := 0; j < len(source[i]); j++ {
			pair := string(source[i][j])
			if j < len(source[i])-1 {
				pair += string(source[i][j+1])
			}

			if blockOpen {
				if pair == "*/" {
					blockOpen = false
					j++
				}
			} else if pair == "/*" {
				blockOpen = true
				j++
			} else if pair == "//" {
				if !blockOpen {
					break
				}
			} else {
				line.WriteByte(pair[0])
			}
		}

		if line.Len() > 0 && !blockOpen {
			clean = append(clean, line.String())
			line = strings.Builder{}
		}
	}

	return clean
}
