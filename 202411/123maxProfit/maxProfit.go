package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}

func maxProfit(prices []int) int {
	n := len(prices)
	var dfs func(i, j int, pre bool) int
	inf := math.MaxInt64 / 10
	// i 表示第几天，j表示交易次数
	// pre 表示当前是否持有

	dfs = func(i, j int, pre bool) int {
		if i < 0 {
			return -inf
		}
		// 只能是两次
		if j > 2 {
			return -inf
		}
		ans := 0
		if pre {
			// 当前持有，那么就可以继续持有，或者卖出
			ans = max(ans, dfs(i-1, j, true), dfs(i-1, j, false)+prices[i])
		} else {
			// 当前没有持有，那么可以继续不持有，也可以买入
			ans = max(ans, dfs(i-1, j, false), dfs(i-1, j+1, true)-prices[i])
		}
		return ans
	}
	ans := dfs(n-1, 0, false)
	return ans
}
