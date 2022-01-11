package perfect_squares

func numSquares(n int) int {
	// get squares < n
	squares := make([]int, 0)

	for i := 1; i*i <= n; i++ {
		squares = append(squares, i*i)
	}

	queue := []int{n}
	steps := 1

	// bfs
	for len(queue) > 0 {
		var tmp []int

		for _, remainder := range queue {
			for _, square := range squares {
				x := remainder - square

				if x == 0 {
					return steps
				} else if x > 0 {
					tmp = append(tmp, x)
				}
			}
		}

		steps++
		queue = tmp
	}

	return 0
}
