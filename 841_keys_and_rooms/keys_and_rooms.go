package keys_and_rooms

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))

	stack := [][]int{rooms[0]}
	visited[0] = true

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, room := range top {
			if !visited[room] {
				stack = append(stack, rooms[room])
				visited[room] = true
			}
		}
	}

	for _, room := range visited {
		if !room {
			return false
		}
	}

	return true
}
