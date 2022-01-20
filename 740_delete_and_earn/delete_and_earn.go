package delete_and_earn

import (
	"sort"
)

func deleteAndEarn(nums []int) int {
	count := make(map[int]int)
	sort.Ints(nums)
	uniq := make([]int, 0)

	for _, num := range nums {
		if _, ok := count[num]; !ok {
			uniq = append(uniq, num)
		}

		count[num] += num
	}

	lastOne, lastTwo := 0, 0
	lastTwo, _ = count[uniq[0]]

	if len(uniq) == 1 {
		return lastTwo
	}

	lastOne, _ = count[uniq[1]]

	if uniq[0] < uniq[1]-1 {
		lastOne += lastTwo
	} else {
		lastOne = max(lastTwo, lastOne)
	}

	for i := 2; i < len(uniq); i++ {
		v, _ := count[uniq[i]]

		if uniq[i-1] == uniq[i]-1 {
			lastOne, lastTwo = max(v+lastTwo, lastOne), lastOne
		} else {
			lastOne, lastTwo = v+lastOne, lastOne
		}
	}

	return max(lastOne, lastTwo)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
