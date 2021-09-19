package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for j, w := range nums[i+1:] {
			if v+w == target {
				return []int{i, j + i + 1}
			}
		}
	}

	return []int{}
}

func twoSumHash(nums []int, target int) []int {
	numbers := map[int]int{}

	for i, v := range nums {
		numbers[v] = i
	}

	for i, v := range nums {
		x := target - v

		if j, ok := numbers[x]; ok && j != i {
			return []int{i, j}
		}
	}

	return []int{}
}

func twoSumHashOnePass(nums []int, target int) []int {
	m := map[int]int{}

	for i, v := range nums {
		complement := target - v

		if j, ok := m[complement]; ok && i != j {
			return []int{i, j}
		}

		m[v] = i
	}

	return []int{}
}

func main() {
	// x := []int{1, 2, 3, 4}
	// fmt.Printf("%v\n", x[1:])

	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))

	fmt.Println(twoSumHash([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSumHash([]int{3, 2, 4}, 6))

	fmt.Println(twoSumHashOnePass([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSumHashOnePass([]int{3, 2, 4}, 6))
	fmt.Println(twoSumHashOnePass([]int{3, 3}, 6))
}
