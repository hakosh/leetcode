package kth_largest_element_in_an_array

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, len(nums)-k, 0, len(nums)-1)
}

func partition(nums []int, low, high int) int {
	pivot := nums[high]
	i := low

	for j := low; j < high; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[high] = nums[high], nums[i]

	return i
}

func quickSelect(nums []int, k, low, high int) int {
	if low == high {
		return nums[low]
	}

	pivot := partition(nums, low, high)

	if pivot == k {
		return nums[pivot]
	} else if pivot < k {
		return quickSelect(nums, k, pivot+1, high)
	} else {
		return quickSelect(nums, k, low, pivot-1)
	}
}
