package replace_elements_with_greatest_element_on_the_right_side

func replaceElements(arr []int) []int {
	maxes := make(map[int]int, len(arr)+1)
	max := 0

	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] > max {
			max = arr[i]
		}

		maxes[i] = max
	}

	maxes[len(arr)] = -1

	for i, _ := range arr {
		arr[i] = maxes[i+1]
	}

	return arr
}
