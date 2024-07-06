package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numWays(439, 315977))
}

func numWays(steps int, n int) int {
	mod := int(math.Pow10(9)) + 7
	// 不加这一步的优化，会超时
	// 同时我们根据「最终回到下标 0 位置」可以推断出，最远到达的位置为 step/2（再远就回不来了）。将最远到达位置与数组最大下标取 min 即可确定维度 step 的范围
	mx := min(steps/2, n-1) + 1
	dp := make([][]int, steps+1)
	for i := range dp {
		dp[i] = make([]int, mx)
	}
	dp[0][0] = 1 // 还没开始走时，就在 index=0处

	for j := 1; j <= steps; j++ {
		for i := 0; i < mx; i++ {
			dp[j][i] += dp[j-1][i]
			if i+1 < mx {
				dp[j][i] += dp[j-1][i+1]
			}
			if i-1 >= 0 {
				dp[j][i] += dp[j-1][i-1]
			}
			dp[j][i] %= mod
		}
	}
	return dp[steps][0] % mod
}
