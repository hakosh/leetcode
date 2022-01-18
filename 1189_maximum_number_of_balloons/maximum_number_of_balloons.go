package maximum_number_of_balloons

func maxNumberOfBalloons(text string) int {
	chars := make([]int, 128)

	for _, char := range text {
		chars[char]++
	}

	answer := 0

	for {
		for _, char := range "balloon" {
			if chars[char] == 0 {
				return answer
			}

			chars[char]--
		}

		answer++
	}
}
