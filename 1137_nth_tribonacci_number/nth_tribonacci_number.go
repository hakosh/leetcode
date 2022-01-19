package nth_tribonacci_number

func tribonacci(n int) int {
	nums := make([]int, 0, n)
	nums = append(nums, 0, 1, 1, 2)

	for i := 4; i <= n; i++ {
		nums = append(nums, nums[i-1]+nums[i-2]+nums[i-3])
	}

	return nums[n]
}
