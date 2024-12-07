package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
}

func maxProfit(prices []int) int {
	n := len(prices)
	if len(prices) == 0 {
		return 0
	}

	dp1 := make([]int, n) // 持有
	dp2 := make([]int, n) // 不持有
	dp3 := make([]int, n) // 冷静

	// 初值
	dp1[0] = -prices[0]

	for i := 1; i < n; i++ {
		// 这里 dp1,dp2,dp3的计算顺序不重要
		dp1[i] = max(dp3[i-1]-prices[i], dp1[i-1])
		dp2[i] = max(dp1[i-1] + prices[i])
		dp3[i] = max(dp3[i-1], dp2[i-1])
	}
	return max(dp2[len(prices)-1], dp3[len(prices)-1])
}
