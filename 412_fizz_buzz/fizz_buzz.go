package fizz_buzz

import "strconv"

func fizzBuzz(n int) []string {
	words := make([]string, 0, n)

	for i := 1; i <= n; i++ {
		this := ""

		if i%3 == 0 {
			this += "Fizz"
		}

		if i%5 == 0 {
			this += "Buzz"
		}

		if this == "" {
			this = strconv.Itoa(i)
		}

		words = append(words, this)
	}

	return words
}
