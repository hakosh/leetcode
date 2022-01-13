package target_sum

func findTargetSumWays(nums []int, target int) int {
	return ways(nums, target)
}

func ways(nums []int, target int) int {
	if len(nums) == 0 {
		if target == 0 {
			return 1
		} else {
			return 0
		}
	}

	num := nums[0]

	return ways(nums[1:], target+num) + ways(nums[1:], target-num)
}
