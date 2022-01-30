package domino_and_tromino_tiling

const mod = 1000000007

var cache [][5]int

func numTilings(n int) int {
	cache = make([][5]int, n+1)
	return dp(n, 0, 1) % mod
}

func dp(n, pos, ending int) int {
	if pos == n {
		switch ending {
		case 1, 4:
			return 1
		case 2, 3:
			return 0
		}
	}

	if pos > n {
		return 0
	}

	if cache[pos][ending] == 0 {
		var ans int

		switch ending {
		case 1:
			ans = dp(n, pos+2, 4) + dp(n, pos+1, 1) + dp(n, pos+2, 3) + dp(n, pos+2, 2)
		case 2:
			ans = dp(n, pos+1, 3) + dp(n, pos+1, 1)
		case 3:
			ans = dp(n, pos+1, 2) + dp(n, pos+1, 1)
		case 4:
			ans = dp(n, pos, 1)
		}

		cache[pos][ending] = ans % mod
	}

	return cache[pos][ending]
}
