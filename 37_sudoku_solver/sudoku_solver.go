package sudoku_solver

var values = [9]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

type cache [58]bool

func solveSudoku(board [][]byte) {
	row := [9]cache{}
	col := [9]cache{}
	box := [9]cache{}

	for x := range board {
		for y, b := range board[x] {
			if b != '.' {
				boxId := getBoxId(x, y)
				add(&box[boxId], &row[x], &col[y], b)
			}
		}
	}

	solve(board, &box, &row, &col, 0, 0)
}

func solve(board [][]byte, box, row, col *[9]cache, x, y int) bool {
	if x == 9 {
		return true
	}

	if board[x][y] == '.' {
		for _, k := range values {
			board[x][y] = k
			boxId := getBoxId(x, y)

			if isValid(&box[boxId], &row[x], &col[y], k) {
				add(&box[boxId], &row[x], &col[y], k)

				nx, ny := next(x, y)
				if solve(board, box, row, col, nx, ny) {
					return true
				}

				del(&box[boxId], &row[x], &col[y], k)
			}

			board[x][y] = '.'
		}

		return false
	} else {
		x, y := next(x, y)
		return solve(board, box, row, col, x, y)
	}
}

func getBoxId(x, y int) int {
	return (x/3)*3 + (y / 3)
}

func add(box, row, col *cache, i byte) {
	box[i] = true
	row[i] = true
	col[i] = true
}

func del(box, row, col *cache, i byte) {
	box[i] = false
	row[i] = false
	col[i] = false
}

func isValid(box, row, col *cache, i byte) bool {
	return !(box[i] || row[i] || col[i])
}

func next(x, y int) (int, int) {
	if y == 8 {
		return x + 1, 0
	}

	return x, y + 1
}
