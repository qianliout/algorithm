package main

import (
	"math"
)

func main() {

}

/*
给你两个下标从 0 开始的字符串 source 和 target ，它们的长度均为 n 并且由 小写 英文字母组成。
另给你两个下标从 0 开始的字符数组 original 和 changed ，以及一个整数数组 cost ，其中 cost[i] 代表将字符 original[i] 更改为字符 changed[i] 的成本。
你从字符串 source 开始。在一次操作中，如果 存在 任意 下标 j 满足 cost[j] == z  、original[j] == x 以及 changed[j] == y 。
你就可以选择字符串中的一个字符 x 并以 z 的成本将其更改为字符 y 。
返回将字符串 source 转换为字符串 target 所需的 最小 成本。如果不可能完成转换，则返回 -1 。
注意，可能存在下标 i 、j 使得 original[j] == original[i] 且 changed[j] == changed[i] 。
*/

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	inf := math.MaxInt / 2
	n := 26
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
		}
	}

	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			// 这里的初值要注意
			if i != j {
				g[i][j] = inf
			}
		}
	}

	for i := 0; i < len(original); i++ {
		x, y, z := original[i]-'a', changed[i]-'a', cost[i]
		g[x][y] = min(g[x][y], z)
	}

	dp[0] = g
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dp[k+1][i][j] = min(dp[k][i][j], dp[k][i][k]+dp[k][k][j])
			}
		}
	}
	ans := 0
	for i := 0; i < len(source); i++ {
		a, b := source[i]-'a', target[i]-'a'
		if dp[n][a][b] >= inf {
			return -1
		}
		ans += dp[n][a][b]
	}
	return int64(ans)
}
