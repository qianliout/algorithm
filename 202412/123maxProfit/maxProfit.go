package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
}

func maxProfit(prices []int) int {
	return maxProfitK(prices, 2)
}

func maxProfitK(prices []int, k int) int {
	n := len(prices)
	has := make([][]int, n)
	notHas := make([][]int, n)
	for i := range has {
		has[i] = make([]int, k+1)
		notHas[i] = make([]int, k+1)
	}
	// 初值
	mx := 0
	for i := 1; i <= k; i++ {
		has[0][i] = -prices[0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			has[i][j] = max(has[i-1][j], notHas[i-1][j-1]-prices[i])
			notHas[i][j] = max(notHas[i-1][j], has[i-1][j]+prices[i])
			mx = max(mx, notHas[i][j])
		}
	}
	return mx
}
