package find_numbers_with_even_number_of_digits

func findNumbers(nums []int) int {
	count := 0

	for _, v := range nums {
		if isEvenDigits(v) {
			count++
		}
	}

	return count
}

func isEvenDigits(num int) bool {
	digits := 0

	for num > 0 {
		digits++
		num /= 10
	}

	return digits%2 == 0
}
