package guess_number_higher_or_lower

var num int

func setNum(pick int) {
	num = pick
}

func guess(pick int) int {
	if num < pick {
		return -1
	} else if num > pick {
		return 1
	} else {
		return 0
	}
}

func guessNumber(n int) int {
	left, right := 0, n

	for {
		pick := (left + right) / 2
		result := guess(pick)

		if result == 0 {
			return pick
		} else if result == -1 {
			right = pick - 1
		} else {
			left = pick + 1
		}
	}
}
