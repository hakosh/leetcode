package boats_to_save_people

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	boats := 0
	left, right := 0, len(people)-1

	for left <= right {
		boats++

		if people[left]+people[right] <= limit {
			left++
		}

		right--
	}

	return boats
}
