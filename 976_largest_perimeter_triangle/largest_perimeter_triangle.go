package largest_perimeter_triangle

import "sort"

func largestPerimeter(nums []int) int {
	sort.Ints(nums)

	p := 0

	for i := len(nums) - 1; i >= 2; i-- {
		a, b, c := nums[i-2], nums[i-1], nums[i]

		if a+b > c {
			p = a + b + c
			break
		}
	}

	return p
}
