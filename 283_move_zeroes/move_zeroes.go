package move_zeroes

func moveZeroes(nums []int) {
	j := 0

	for i, num := range nums {
		if num == 0 {
			for _, num := range nums[j:] {
				if num != 0 {
					break
				}

				j++
			}

			if j >= len(nums) {
				break
			}

			nums[i], nums[j] = nums[j], 0
		}

		j++
	}
}
