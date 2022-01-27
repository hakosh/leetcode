package decode_ways

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1

	if s[0] != '0' {
		dp[1] = 1
	}

	for i := 2; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		v, _ := strconv.Atoi(s[i-2 : i])

		if v >= 10 && v <= 26 {
			dp[i] += dp[i-2]
		}
	}

	fmt.Println(dp)

	return dp[len(s)]
}
