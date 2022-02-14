package find_n_unique_integers_sum_up_to_zero

func sumZero(n int) []int {
	nums := make([]int, 0, n)

	for i := 1; i <= n/2; i++ {
		nums = append(nums, i, -i)
	}

	if n%2 == 1 {
		nums = append(nums, 0)
	}

	return nums
}
