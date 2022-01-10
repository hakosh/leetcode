package walls_and_gates

import "math"

type Room struct {
	row int
	col int
}

func wallsAndGates(rooms [][]int) {
	// Find gates
	var queue []Room

	m := len(rooms)
	n := len(rooms[0])

	for i, row := range rooms {
		for j, col := range row {
			if col == 0 {
				queue = append(queue, Room{i, j})
			}
		}
	}

	// BFS for each gate
	step := 0

	for len(queue) > 0 {
		var tmp []Room
		step++

		for _, room := range queue {
			neighbors := []Room{
				{max(0, room.row-1), room.col},
				{min(m-1, room.row+1), room.col},
				{room.row, max(0, room.col-1)},
				{room.row, min(n-1, room.col+1)},
			}

			for _, room := range neighbors {
				if rooms[room.row][room.col] == math.MaxInt32 {
					rooms[room.row][room.col] = step
					tmp = append(tmp, Room{room.row, room.col})
				}
			}
		}

		queue = tmp
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
