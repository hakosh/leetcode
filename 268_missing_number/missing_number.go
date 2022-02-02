package missing_number

func missingNumber(nums []int) int {
	sum1 := len(nums)
	sum2 := 0

	for i := 0; i < len(nums); i++ {
		sum1 += i
		sum2 += nums[i]
	}

	return sum1 - sum2
}
