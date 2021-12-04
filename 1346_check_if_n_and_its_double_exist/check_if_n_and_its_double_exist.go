package check_if_n_and_its_double_exist

func checkIfExist(arr []int) bool {
	seen := make(map[int]bool, len(arr))

	for _, v := range arr {
		if _, ok := seen[v*2]; ok {
			return true
		}

		if v%2 == 0 {
			if _, ok := seen[v/2]; ok {
				return true
			}
		}

		seen[v] = true
	}

	return false
}
