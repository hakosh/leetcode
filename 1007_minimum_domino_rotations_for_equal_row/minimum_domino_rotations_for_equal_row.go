package minimum_domino_rotations_for_equal_row

func minDominoRotations(tops []int, bottoms []int) int {
	if a := count(tops, bottoms, tops[0]); a != -1 || tops[0] == bottoms[0] {
		return a
	} else {
		return count(tops, bottoms, bottoms[0])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func count(tops []int, bottoms []int, target int) int {
	a, b := 0, 0

	for i := 0; i < len(tops); i++ {
		if tops[i] != target && bottoms[i] != target {
			return -1
		} else if tops[i] != target {
			a++
		} else if bottoms[i] != target {
			b++
		}
	}

	return min(a, b)
}
