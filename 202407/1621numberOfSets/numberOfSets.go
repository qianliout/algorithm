package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numberOfSets(4, 2))
}

func numberOfSets(n int, k int) int {
	mod := int(math.Pow10(9)) + 7
	// dp[i][j]表示 前i 个点组合 k 个线段
	dp1 := make([][]int, n) // 表示第 j 条线段的右端点不是 i
	dp2 := make([][]int, n) //  表示第 j 条线段的右端点就是 i
	for i := 0; i < n; i++ {
		dp1[i] = make([]int, k+1)
		dp2[i] = make([]int, k+1)
	}

	// 初值,只能写这个， 是为什么呢
	dp1[0][0] = 1
	// dp2[0][0] = 1 不能写
	for i := 1; i < n; i++ {
		for j := 0; j <= k; j++ {
			// 首先考虑 dp1[i][j]，因为第 j 条线段的右端点不是 i，因此第 i 个点没有用上，那么 0 .. i-1 的点构造了 j 条线段，即
			dp1[i][j] += dp1[i-1][j] + dp2[i-1][j]
			dp1[i][j] %= mod
			// 再考虑 dp2[i][j]，因为第 j 条线段的右端点就是 i，因此有两种情况：
			//    第 j 条线段长度为 1，那么 0 .. i-1 的点构造了 j−1 条线段，即

			if j > 0 {
				dp2[i][j] += dp1[i-1][j-1] + dp2[i-1][j-1]
				dp2[i][j] %= mod
			}

			// 第 j 条线段长度大于 1，那么删去第 j 条线段 i-1 .. i 的这一部分，0 .. i-1 的点仍然构造了 j 条线段，并且点 i−1 是属于第 j 条线段的，即
			dp2[i][j] += dp2[i-1][j]
			dp2[i][j] %= mod
		}
	}
	return (dp1[n-1][k] + dp2[n-1][k]) % mod
}
