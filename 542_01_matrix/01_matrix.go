package matrix

import "math"

func updateMatrix(mat [][]int) [][]int {
	var queue [][2]int

	m := len(mat)
	n := len(mat[0])

	dist := make([][]int, m)

	for i, row := range mat {
		dist[i] = make([]int, n)

		for j, col := range row {
			if col == 0 {
				queue = append(queue, [2]int{i, j})
				dist[i][j] = 0
			} else {
				dist[i][j] = math.MaxInt32
			}
		}
	}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		neighbors := [][2]int{
			{p[0] - 1, p[1]},
			{p[0] + 1, p[1]},
			{p[0], p[1] - 1},
			{p[0], p[1] + 1},
		}

		for _, nbr := range neighbors {
			if nbr[0] < 0 || nbr[0] >= m || nbr[1] < 0 || nbr[1] >= n {
				continue
			}

			if dist[nbr[0]][nbr[1]] > dist[p[0]][p[1]]+1 {
				dist[nbr[0]][nbr[1]] = dist[p[0]][p[1]] + 1
				queue = append(queue, nbr)
			}
		}

	}

	return dist
}
