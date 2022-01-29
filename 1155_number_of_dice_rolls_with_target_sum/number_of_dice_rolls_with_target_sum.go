package number_of_dice_rolls_with_target_sum

import "fmt"

const mod = 1000000007

// BOTTOM UP

func numRollsToTarget(n, k, target int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}

	for i := 1; i <= k; i++ {
		if i <= target {
			dp[1][i] = 1
		}
	}

	for i := 1; i <= n; i++ {
		for t := 1; t <= target; t++ {
			for s := 1; s <= k; s++ {
				if t >= s {
					dp[i][t] = (dp[i][t] + dp[i-1][t-s]) % mod
				}
			}
		}
	}

	for i := range dp {
		fmt.Println(dp[i])
	}

	return dp[n][target]
}

// TOP DOWN

var cache [][]int

func numRollsToTargetR(n int, k int, target int) int {
	cache = make([][]int, n+1)
	for i := range cache {
		cache[i] = make([]int, target+1)
	}

	return dp(n, k, target) % mod
}

func dp(dice, sides, left int) int {
	if dice*sides < left || left < 0 || (left == 0 && dice > 1) {
		return 0
	}

	if dice == 0 {
		if left == 0 {
			return 1
		} else {
			return 0
		}
	}

	if cache[dice][left] == 0 {
		var ans int

		for i := sides; i > 0; i-- {
			sub := dp(dice-1, sides, left-i)
			ans += sub % mod
		}

		cache[dice][left] = ans
	}

	return cache[dice][left]
}
