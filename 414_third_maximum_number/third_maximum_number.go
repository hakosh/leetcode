package third_maximum_number

// Checking if the elements are unique is messy, but works on 32bit

func thirdMax(nums []int) int {
	max1, max2, max3 := nums[0], nums[0], nums[0]

	for _, v := range nums {
		if v > max1 {
			max1, max2, max3 = v, max1, max2
		} else if max2 == max1 || (v > max2 && v < max1) {
			max2, max3 = v, max2
		} else if max3 == max2 || max3 == max1 || (v > max3 && v < max2) {
			max3 = v
		}
	}

	if max1 == max2 || max2 == max3 {
		return max1
	}

	return max3
}
