package subarray_sum_equals_k

func subarraySum(nums []int, k int) int {
	freq := make(map[int]int)
	freq[0] = 1

	count, sum := 0, 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		v, ok := freq[sum-k]

		if ok {
			count += v
		}

		freq[sum]++
	}

	return count
}
