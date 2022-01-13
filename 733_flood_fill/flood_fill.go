package flood_fill

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	queue := [][2]int{{sr, sc}}
	color := image[sr][sc]

	if color == newColor {
		return image
	}

	m := len(image)
	n := len(image[0])

	for len(queue) > 0 {
		tmp := make([][2]int, 0)

		for _, p := range queue {
			image[p[0]][p[1]] = newColor

			neighbors := [][2]int{
				{p[0] - 1, p[1]},
				{p[0] + 1, p[1]},
				{p[0], p[1] - 1},
				{p[0], p[1] + 1},
			}

			for _, nbr := range neighbors {
				if nbr[0] < 0 || nbr[0] >= m || nbr[1] < 0 || nbr[1] >= n || image[nbr[0]][nbr[1]] != color {
					continue
				}

				tmp = append(tmp, nbr)
			}
		}

		queue = tmp
	}

	return image
}
