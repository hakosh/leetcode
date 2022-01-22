package single_number

func singleNumber(nums []int) int {
	n := 0

	for _, num := range nums {
		n ^= num
	}

	return n
}

func singleNumberSet(nums []int) int {
	m := make(map[int]bool)

	for _, num := range nums {
		if _, ok := m[num]; ok {
			delete(m, num)
		} else {
			m[num] = true
		}
	}

	num := 0

	for k, _ := range m {
		num = k
	}

	return num
}
