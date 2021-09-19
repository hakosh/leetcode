package longest_common_prefix

// longestCommonPrefix returns the longest common prefix
// amongst an array of string. If there is no common prefix,
// it returns an empty string.
func longestCommonPrefix(strs []string) string {
	prefix := ""

	if len(strs[0]) == 0 {
		return prefix
	}

	for _, letter := range strs[0] {
		test := prefix + string(letter)

		for _, word := range strs {
			if len(prefix) == len(word) {
				return prefix
			}

			if test != word[0:len(test)] {
				return prefix
			}
		}

		prefix = test
	}

	return prefix
}
