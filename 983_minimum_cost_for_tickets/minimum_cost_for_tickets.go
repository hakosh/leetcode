package minimum_cost_for_tickets

var cache [][395]int

func mincostTickets(days []int, costs []int) int {
	cache = make([][395]int, len(days))
	return dp(days, costs, 0, 0)
}

func dp(days, costs []int, day, until int) int {
	if day == len(days) {
		return 0
	}

	if cache[day][until] == 0 {
		var ans int

		if until > days[day] {
			ans = dp(days, costs, day+1, until)
		} else {
			ans = min(
				costs[0]+dp(days, costs, day+1, days[day]+1),
				min(
					costs[1]+dp(days, costs, day+1, days[day]+7),
					costs[2]+dp(days, costs, day+1, days[day]+30),
				),
			)
		}

		cache[day][until] = ans
	}

	return cache[day][until]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
