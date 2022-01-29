package maximum_length_of_repeated_subarray

func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	best := 0

	for i := len(nums1) - 1; i >= 0; i-- {
		for j := len(nums2) - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			}

			if dp[i][j] > best {
				best = dp[i][j]
			}
		}
	}

	return best
}
