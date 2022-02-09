package course_schedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses == 0 {
		return true
	}

	follows := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	for _, pre := range prerequisites {
		follows[pre[1]] = append(follows[pre[1]], pre[0])
		inDegree[pre[0]]++
	}

	start := make([]int, 0)
	for i, l := range inDegree {
		if l == 0 {
			start = append(start, i)
		}
	}

	return bfs(follows, inDegree, start)
}

func bfs(follows [][]int, inDegree []int, start []int) bool {
	n := len(follows)
	visited := make([]bool, n)

	queue := start

	for len(queue) > 0 {
		tmp := make([]int, 0)

		for _, class := range queue {
			if visited[class] {
				continue
			}

			visited[class] = true
			n--

			for _, class := range follows[class] {
				inDegree[class]--

				if inDegree[class] == 0 {
					tmp = append(tmp, class)
				}
			}
		}

		queue = tmp
	}

	return n == 0
}
