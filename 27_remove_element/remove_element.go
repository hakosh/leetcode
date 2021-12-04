package remove_element

func removeElement(nums []int, val int) int {
	shifts := 0

	ri := len(nums) - 1

	for li := 0; li <= ri; li++ {
		if nums[li] == val {
			shifts++

			for {
				if nums[ri] != val || ri <= li {
					break
				} else {
					shifts++
					nums[ri] = 0
					ri--
				}
			}

			nums[li], nums[ri] = nums[ri], 0

			ri--
		}
	}

	return len(nums) - shifts
}
