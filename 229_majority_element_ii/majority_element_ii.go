package majority_element_ii

import (
	"math"
	"sort"
)

// Boyer-Moore Voting Algorithm
func majorityElement(nums []int) []int {
	candidate1, candidate2 := math.MinInt32, math.MinInt32
	count1, count2 := 0, 0

	for _, num := range nums {
		if candidate1 == num {
			count1++
		} else if candidate2 == num {
			count2++
		} else if count1 == 0 {
			candidate1 = num
			count1++
		} else if count2 == 0 {
			candidate2 = num
			count2++
		} else {
			count1--
			count2--
		}
	}

	matches := make([]int, 0, 2)
	count1, count2 = 0, 0

	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		}
	}

	if count1 > len(nums)/3 {
		matches = append(matches, candidate1)
	}

	if count2 > len(nums)/3 {
		matches = append(matches, candidate2)
	}

	return matches
}

func majorityElementSort(nums []int) []int {
	sort.Ints(nums)
	current := 0
	threshold := len(nums) / 3

	matches := make([]int, 0, 3)
	for i := 0; i <= len(nums); i++ {
		if i == len(nums) || nums[i] != nums[current] {
			if i-current > threshold {
				matches = append(matches, nums[current])
			}

			current = i
		}
	}

	return matches
}
