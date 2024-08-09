package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(longestIdealString("acfgbd", 2))
}

func longestIdealString(s string, k int) int {
	// dp[i][c] 前定义 f[i][c] 表示 s 的前 i 个字母中的以 c 结尾的理想字符串的最长长度。
	dp := make([]int, 26)
	for _, ch := range s {
		c := int(ch) - int('a')
		mx := 0
		for j := c - k; j <= c+k; j++ {
			if j < 0 || j >= 26 {
				continue
			}
			mx = max(mx, dp[j]+1)
		}
		dp[c] = max(mx, dp[c])
	}
	return slices.Max(dp)
}
