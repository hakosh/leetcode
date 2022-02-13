package maximum_length_of_a_concatenated_string_with_unique_characters

func maxLength(arr []string) int {
	return solve(make([]int, 127), arr, 0)
}

func solve(chars []int, arr []string, current int) int {
	if len(arr) == 0 {
		return current
	}

	max := current

	for i, s := range arr {
		if concatenate(chars, s) {
			length := solve(chars, arr[i+1:], current+len(s))

			if length > max {
				max = length
			}

			remove(chars, s)
		}
	}

	return max
}

func remove(chars []int, str string) {
	for _, c := range str {
		chars[c]--
	}
}

func concatenate(chars []int, str string) bool {
	added := 0

	for i, c := range str {
		if chars[c] == 0 {
			chars[c]++
			added++
		} else {
			for _, c := range str[:i] {
				chars[c]--
				added--
			}

			break
		}
	}

	return added > 0
}
