package main

import (
	"math"
)

func main() {

}

func countHousePlacements(n int) int {
	mod := int(math.Pow10(9)) + 7
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i <= n; i++ {
		// dp[i] = dp[i-1]           // 不放 i,那么第 i−1 个地块可放可不放，则有 f[i]=f[i−1]；
		dp[i] = dp[i-1] + dp[i-2] // 放置第i，那么第 i−1 个地块无法放房子，第 i−2 个地块可放可不放，则有 f[i]=f[i−2]
		dp[i] = dp[i] % mod
	}
	return (dp[n] * dp[n]) % mod
}
