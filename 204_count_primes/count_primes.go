package count_primes

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func countPrimes(n int) int {
	nums := make([]bool, n/2)

	count := 0

	for i := 3; i*i <= n; i++ {
		if isPrime(i) {
			for j, k := i*3, 3; j < n; k, j = k+2, i*k {
				nums[j/2] = true
			}
		}
	}

	if n > 2 {
		count++
	}

	for i := 1; i < len(nums); i++ {
		if !nums[i] {
			count++
		}
	}

	return count
}
