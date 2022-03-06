package count_all_valid_pickup_and_delivery_options

const MOD = 1000000007

var cache [][]int

func countOrders(n int) int {
	cache = make([][]int, n+1)
	for i := range cache {
		cache[i] = make([]int, n+1)
	}

	return dp(n, n)
}

func dp(unpicked, undelivered int) int {
	if unpicked == 0 && undelivered == 0 {
		return 1
	}

	if unpicked < 0 || undelivered < 0 || undelivered < unpicked {
		return 0
	}

	if cache[unpicked][undelivered] == 0 {
		up := unpicked * dp(unpicked-1, undelivered) % MOD
		ud := (undelivered - unpicked) * dp(unpicked, undelivered-1) % MOD

		cache[unpicked][undelivered] = up + ud
	}

	return cache[unpicked][undelivered]
}
