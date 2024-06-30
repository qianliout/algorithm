package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumTimeRequired([]int{3, 2, 3}, 3))
}

func minimumTimeRequired(jobs []int, k int) int {
	n := len(jobs)
	m := 1 << n
	sum := make([]int, m)
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i>>j)&1 == 1 {
				sum[i] += jobs[j]
			}
		}
	}

	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, m)
	}
	inf := math.MaxInt / 10
	dp[1] = sum // 分配置一个人
	for i := 2; i <= k; i++ {
		for j := 0; j < m; j++ {
			dp[i][j] = inf
			for s := j; s > 0; s = (s - 1) & j {
				dp[i][j] = min(dp[i][j], max(dp[i-1][j^s], sum[s]))
			}
		}
	}
	return dp[k][m-1]
}
