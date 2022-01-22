package rotate_array

func rotate(nums []int, k int) {
	if len(nums) < 2 {
		return
	}

	k = k % len(nums)
	if k == 0 {
		return
	}

	offset := len(nums) - k
	for i := 0; i < k; i++ {
		nums[i], nums[offset+i] = nums[offset+i], nums[i]
	}

	rotate(nums[k:], k)
}
