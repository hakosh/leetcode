package counting_bits

func countBits(n int) []int {
	ans := make([]int, n+1)

	for i := 1; i <= n; i++ {
		ans[i] = ans[i&(i-1)] + 1
	}

	return ans
}

func countBitsIterative(n int) []int {
	ans := make([]int, n+1)

	for i := 0; i <= n; i++ {
		ans[i] = hammingWeight(i)
	}

	return ans
}

func hammingWeight(n int) int {
	weight := 0

	for n > 0 {
		if n%2 == 1 {
			weight++
		}

		n = n >> 1
	}

	return weight
}
