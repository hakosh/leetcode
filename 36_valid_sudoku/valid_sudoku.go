package valid_sudoku

func isValidSudoku(board [][]byte) bool {
	for i := range board {
		numsRow := make([]uint8, 9)
		numsCol := make([]uint8, 9)

		for j := range board {
			valRow := board[i][j]
			valCol := board[j][i]

			if valRow != '.' {
				val := valRow - 49

				if numsRow[val] != 0 {
					return false
				} else {
					numsRow[val]++
				}
			}

			if valCol != '.' {
				val := valCol - 49

				if numsCol[val] != 0 {
					return false
				} else {
					numsCol[val]++
				}
			}
		}
	}

	// Check subgrids
	for offsetX := 0; offsetX < 7; offsetX += 3 {
		for offsetY := 0; offsetY < 7; offsetY += 3 {
			nums := make([]uint8, 9)

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					val := board[offsetX+i][offsetY+j]

					if val == '.' {
						continue
					}

					if nums[val-49] != 0 {
						return false
					} else {
						nums[val-49]++
					}
				}
			}
		}
	}

	return true
}
