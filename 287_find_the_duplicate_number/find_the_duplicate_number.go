package find_the_duplicate_number

func findDuplicate(nums []int) int {
	duplicate := 0
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		count := 0

		for _, num := range nums {
			if num <= mid {
				count++
			}
		}

		if count > mid {
			duplicate = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return duplicate
}
