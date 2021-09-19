package string_to_integer

import (
	"math"
)

func myAtoi(s string) int {
	for _, char := range s {
		if char == ' ' {
			s = s[1:]
		} else {
			break
		}
	}

	if len(s) == 0 {
		return 0
	}

	result := 0
	from := 0
	sign := 1
	overflow := false

	if s[0] == '-' {
		sign = -1
		from = 1
	} else if s[0] == '+' {
		from = 1
	}

	for _, char := range s[from:] {
		if char < '0' || char > '9' {
			break
		}

		if result == 0 && char == '0' {
			continue
		}

		base := result * 10
		if base / 10 != result {
			overflow = true
			break
		}

		result = base + (int(char) - 48)

		if result > math.MaxInt32 {
			overflow = true
			break
		}
	}

	if overflow && sign == -1 {
		return math.MinInt32
	} else if overflow {
		return math.MaxInt32
	} else if sign == -1 {
		return -result
	} else {
		return result
	}
}