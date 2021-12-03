package duplicate_zeros

func duplicateZeros(arr []int) {
	tmp := make([]int, len(arr))
	idx := 0

	for i := 0; idx < len(arr); i++ {
		if arr[i] == 0 && idx+1 != len(arr) {
			tmp[idx] = 0
			tmp[idx+1] = 0

			idx += 2
		} else {
			tmp[idx] = arr[i]

			idx += 1
		}
	}

	copy(arr, tmp)
}
