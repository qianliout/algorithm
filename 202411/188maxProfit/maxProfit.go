package main

import (
	"strings"
)

func main() {

}

func maxProfit(k int, prices []int) int {
	n := len(prices)
	k = k % len(prices)
	if k == 0 {
		k = len(prices)
	}
	has := make([][]int, n)
	notHas := make([][]int, n)
	for i := range has {
		has[i] = make([]int, k+1)
		notHas[i] = make([]int, k+1)
	}
	// 初值
	for i := 0; i <= k; i++ {
		notHas[0][i] = 0
		has[0][i] = -prices[0]
	}

	ans := 0
	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			has[i][j] = max(has[i-1][j], notHas[i-1][j-1]-prices[i])
			notHas[i][j] = max(notHas[i-1][j], has[i-1][j]+prices[i])
			ans = max(ans, notHas[i][j])
		}
	}

	return ans
}

// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k笔 交易。

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
