package the_k_weakest_row_in_a_matrix

import (
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {
	rows := make([][2]int, len(mat))
	for i, row := range mat {
		rows[i] = [2]int{i, strength(row)}
	}

	sort.SliceStable(rows, func(i, j int) bool {
		if rows[i][1] == rows[j][1] {
			return i < j
		}

		return rows[i][1] < rows[j][1]
	})

	ans := make([]int, k)
	for i := range ans {
		ans[i] = rows[i][0]
	}

	return ans
}

func strength(row []int) int {
	return sort.Search(len(row), func(i int) bool {
		return row[i] <= 0
	})
}
