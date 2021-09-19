package _3_roman_to_integer

var values = map[rune]int {
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var (
		result int
		last rune
	)

	for _, char := range s {
		value, _ := values[char]

		if last == 'I' && (char == 'V' || char == 'X') {
			value -= 2
		} else if last == 'X' && (char == 'L' || char == 'C') {
			value -= 20
		} else if last == 'C' && (char == 'D' || char == 'M') {
			value -= 200
		}

		result += value
		last = char
	}

	return result
}
