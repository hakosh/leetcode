package fibonacci_number

var m = make(map[int]int)

func fib(n int) int {
	if n == 1 || n == 0 {
		return n
	}

	if v, ok := m[n]; ok {
		return v
	}

	r := fib(n-1) + fib(n-2)
	m[n] = r

	return r
}
