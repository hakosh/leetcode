package sort_integers_by_the_number_of_1_bits

import (
	"math/bits"
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		a := bits.OnesCount32(uint32(arr[i]))
		b := bits.OnesCount32(uint32(arr[j]))

		if a == b {
			return arr[i] < arr[j]
		} else {
			return a < b
		}
	})

	return arr
}
