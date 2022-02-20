package minimum_size_subarray

func minSubArrayLen(target int, nums []int) int {
	leftSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		leftSum[i+1] = nums[i] + leftSum[i]
	}

	if leftSum[len(leftSum)-1] < target {
		return 0
	}

	size := 0
	left, right := 0, 1

	for right < len(leftSum) {
		if leftSum[right]-leftSum[left] < target {
			right++
		} else {
			if size == 0 || right-left < size {
				size = right - left
			}

			left++
		}
	}

	return size
}

func minSubArrayLenIterative(target int, nums []int) int {
	size := 0

	for i := 0; i < len(nums); i++ {
		sum := target

		for j := i; j < len(nums); j++ {
			sum -= nums[j]

			if sum <= 0 && (size == 0 || size > j-i+1) {
				size = j - i + 1
				break
			}
		}
	}

	return size
}
