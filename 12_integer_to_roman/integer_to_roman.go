package integer_to_roman

var values = []struct{
	symbol string
	value int
}{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

func integerToRoman(num int) string {
	roman := ""
	idx := 0

	for num > 0 {
		if num >= values[idx].value {
			roman += values[idx].symbol
			num -= values[idx].value
		} else {
			idx++
		}
	}

	return roman
}
