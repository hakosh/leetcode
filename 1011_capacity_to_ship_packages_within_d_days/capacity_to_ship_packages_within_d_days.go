package capacity_to_ship_packages_within_d_days

import "sort"

func shipWithinDays(weights []int, days int) int {
	max := 0
	sum := 0

	for _, weight := range weights {
		sum += weight

		if weight > max {
			max = weight
		}
	}

	return max + sort.Search(sum-max+1, func(i int) bool {
		return daysRequired(weights, i+max) <= days
	})

	//low, high := max, sum
	//best := 0
	//
	//for low <= high {
	//	mid := (low + high) / 2
	//
	//	if daysRequired(weights, mid) <= days {
	//		high = mid - 1
	//		best = mid
	//	} else {
	//		low = mid + 1
	//	}
	//}
	//
	//return best
}

func daysRequired(weights []int, capacity int) int {
	days := 0
	this := capacity

	for _, weight := range weights {
		if this+weight <= capacity {
			this += weight
		} else {
			this = weight
			days++
		}
	}

	return days
}
