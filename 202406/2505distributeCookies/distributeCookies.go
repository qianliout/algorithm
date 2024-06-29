package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(distributeCookies([]int{8, 15, 10, 20, 8}, 2))
}

func distributeCookies(cookies []int, k int) int {
	n := len(cookies)
	m := 1 << n

	sum := make([]int, m)
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i>>j)&1 == 1 {
				sum[i] += cookies[j]
			}
		}
	}
	dp := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	inf := math.MaxInt / 10
	// 初值
	dp[0] = sum // 没有能理解
	for i := 1; i < k; i++ {
		for j := 0; j < m; j++ {
			dp[i][j] = inf
			for s := j; s > 0; s = (s - 1) & j {
				dp[i][j] = min(dp[i][j], max(dp[i-1][j^s], sum[s]))
			}
		}
	}
	return dp[k-1][m-1]
}
