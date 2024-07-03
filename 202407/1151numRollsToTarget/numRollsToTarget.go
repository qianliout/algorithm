package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numRollsToTarget(30, 30, 500))
}

func numRollsToTarget(n int, k int, target int) int {
	mod := int(math.Pow10(9)) + 7
	// dp[i][j] 表示执i 次骰子，得到结果j 的结果集数
	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, target+1)
	}
	// 初值
	for j := 1; j <= min(target, k); j++ {
		dp[1][j] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			for m := 1; j-m >= 1 && m <= k; m++ {
				dp[i][j] = (dp[i][j] + dp[i-1][j-m]) % mod
			}
		}
	}
	return dp[n][target] % mod
}
