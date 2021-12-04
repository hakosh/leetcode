package remove_duplicates

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	k := 0

	for i := 1; i < len(nums); i++ {
		nums[k+1] = nums[i]

		if nums[k] < nums[k+1] {
			k++
		}
	}

	return k + 1
}
