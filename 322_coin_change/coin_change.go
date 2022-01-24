package coin_change

import (
	"math"
	"sort"
)

// BOTTOM UP

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}

	dp[0] = 0

	for i := 0; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = minFn(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] <= amount {
		return dp[amount]
	} else {
		return -1
	}
}

func minFn(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// TOP DOWN

var min = math.MaxInt32
var cache [10000]int

func coinChangeR(coins []int, amount int) int {
	for i := range cache {
		cache[i] = -2
	}

	sort.Ints(coins)
	min = coins[0]

	return dp(coins, amount)
}

func dp(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	if amount < min {
		return -1
	}

	if v := cache[amount]; v != -2 {
		return v
	}

	best := -2

	for _, coin := range coins {
		res := dp(coins, amount-coin)

		if res == -1 {
			continue
		} else if best == -2 || res < best {
			best = res
		}
	}

	cache[amount] = best + 1

	return best + 1
}
