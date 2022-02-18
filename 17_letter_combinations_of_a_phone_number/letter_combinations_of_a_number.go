package letter_combinations_of_a_phone_number

var symbols = [][]byte{
	{},
	{},
	{'a', 'b', 'c'},
	{'d', 'e', 'f'},
	{'g', 'h', 'i'},
	{'j', 'k', 'l'},
	{'m', 'n', 'o'},
	{'p', 'q', 'r', 's'},
	{'t', 'u', 'v'},
	{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	combinations := make([]string, 0, 256)

	if len(digits) == 0 {
		return combinations
	}

	var generate func(string, []byte)
	generate = func(digits string, parts []byte) {
		if len(digits) == 0 {
			combinations = append(combinations, string(parts))
			return
		}

		for _, char := range symbols[digits[0]-'0'] {
			this := make([]byte, len(parts), cap(parts))
			copy(this, parts)
			this = append(this, char)
			generate(digits[1:], this)
		}
	}

	generate(digits, make([]byte, 0, len(digits)))

	return combinations
}
