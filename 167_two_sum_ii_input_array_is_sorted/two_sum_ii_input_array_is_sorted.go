package two_sum_ii_input_array_is_sorted

import (
	"sort"
)

func twoSum(numbers []int, target int) []int {
	low, high := 0, len(numbers)-1

	for low < high {
		sum := numbers[low] + numbers[high]

		if sum == target {
			break
		} else if sum > target {
			high--
		} else {
			low++
		}
	}

	return []int{low + 1, high + 1}
}

func twoSumBruteForce(numbers []int, target int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		for j := len(numbers) - 1; j > i; j-- {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}

			if numbers[j] < target-numbers[i] {
				break
			}
		}
	}

	return []int{}
}

func twoSumBinarySearch(numbers []int, target int) []int {
	var pair []int

	for i := 0; i < len(numbers)-1; i++ {
		rightSide := numbers[i+1:]
		complement := target - numbers[i]

		pos := sort.Search(len(rightSide), func(i int) bool {
			return rightSide[i] >= complement
		})

		if pos < len(rightSide) && rightSide[pos] == complement {
			pair = []int{i + 1, i + 2 + pos}
			break
		}
	}

	return pair
}
