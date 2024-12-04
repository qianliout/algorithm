package main

import (
	"fmt"
)

func main() {
	fmt.Println(climbStairs(2))
}

func climbStairs(n int) int {
	dp := make([]int, n+3)
	dp[0] = 1
	for i := 0; i <= n; i++ {
		dp[i+2] = dp[i+1] + dp[i]
	}
	return dp[n+2]
}
