package basic_calculator_ii

func calculate(s string) int {
	op := '+'
	lastNum, thisNum, result := 0, 0, 0

	for i, c := range s {
		if c >= '0' && c <= '9' {
			thisNum = thisNum*10 + int(c-'0')
		}

		if ((c < '0' || c > '9') && c != ' ') || i == len(s)-1 {
			switch op {
			case '+':
				result += lastNum
				lastNum = thisNum
			case '-':
				result += lastNum
				lastNum = -thisNum
			case '*':
				lastNum *= thisNum
			case '/':
				lastNum /= thisNum
			}

			op = c
			thisNum = 0
		}
	}

	return result + lastNum
}
