package find_three_consecutive_integers_that_sum_to_a_given_number

func sumOfThree(num int64) []int64 {
	if num%3 == 0 {
		mid := num / 3
		return []int64{mid - 1, mid, mid + 1}
	}

	return []int64{}
}
