package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	fmt.Println(minSideJumps([]int{0, 1, 2, 3, 0}))
}

// 点 0 处和点 n 处的任一跑道都不会有障碍。
func minSideJumps(obstacles []int) int {
	inf := math.MaxInt64 / 1000
	n := len(obstacles)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 3)

		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	// 初值
	dp[0][0], dp[0][1], dp[0][2] = 1, 0, 1
	for i := 1; i < n; i++ {
		minCnt := inf
		for j := 0; j <= 2; j++ {
			if j != obstacles[i]-1 {
				minCnt = min(minCnt, dp[i-1][j])
			}
		}
		for j := 0; j < 3; j++ {
			if j != obstacles[i]-1 {
				dp[i][j] = min(dp[i-1][j], minCnt+1)
			}
		}
	}
	return slices.Min(dp[n-1])
}
