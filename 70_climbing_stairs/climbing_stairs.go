package climbing_stairs

var mem []int

func climbStairs(n int) int {
	mem = make([]int, n+1)
	return climb(n)
}

func climb(n int) int {
	if n <= 2 {
		return n
	}

	if v := mem[n]; v == 0 {
		mem[n] = climb(n-1) + climb(n-2)
		return mem[n]
	} else {
		return v
	}
}
