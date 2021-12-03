package merge_sorted_array

func merge(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, m+n)

	j, k := 0, 0

	for i := 0; i < len(sorted); i++ {
		if k == n {
			sorted[i] = nums1[j]
			j++
		} else if j == m {
			sorted[i] = nums2[k]
			k++
		} else if nums1[j] < nums2[k] {
			sorted[i] = nums1[j]
			j++
		} else {
			sorted[i] = nums2[k]
			k++
		}
	}

	copy(nums1, sorted)
}
