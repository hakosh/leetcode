package contains_duplicate

func containsDuplicate(nums []int) bool {
	found := make(map[int]bool, len(nums))

	for _, num := range nums {
			if _, ok := found[num]; ok {
				return true
			}

			found[num] = true
	}

	return false
}
