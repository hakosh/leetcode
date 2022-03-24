package boats_to_save_people

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	left, right := 0, len(people)-1
	boats, current, count := 0, 0, 0

	for left-right != 1 {
		if count == 2 {
			boats, current, count = boats+1, 0, 0
		}

		if current+people[right] <= limit {
			current += people[right]
			count++
			right--
		} else if current+people[left] <= limit {
			current += people[left]
			count++
			left++
		} else {
			boats, current, count = boats+1, 0, 0
		}
	}

	if current != 0 {
		boats++
	}

	return boats
}
