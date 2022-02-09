package time_needed_to_inform_all_employees

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	m := make([][]int, n)

	for id, bossId := range manager {
		if bossId == -1 {
			continue
		}

		m[bossId] = append(m[bossId], id)
	}

	return dfs(m, informTime, headID)
}

func dfs(employees [][]int, informTime []int, head int) int {
	subs := employees[head]
	minutes := 0

	if len(subs) == 0 {
		return 0
	}

	for _, sub := range subs {
		minutes = max(minutes, dfs(employees, informTime, sub))
	}

	return informTime[head] + minutes
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
