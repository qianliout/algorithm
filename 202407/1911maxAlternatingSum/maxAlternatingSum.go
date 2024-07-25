package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxAlternatingSum([]int{4, 2, 5, 3}))
	fmt.Println(maxAlternatingSum([]int{5, 6, 7, 8}))
}

func maxAlternatingSum(nums []int) int64 {
	n := len(nums)
	dp1 := make([]int, n+1) // dp[i]表示前 i 个数，且数组长度是偶数
	dp2 := make([]int, n+1) // dp[i]表示前i 个数，且数组长度是奇数数
	dp1[0] = 0
	dp2[0] = math.MinInt / 100

	for i := 0; i < n; i++ {
		// 对于 dp1[i+1] 若不选第 i 个数，则从 dp1[i] 转移过来，否则从 dp2[i]-nums[i] 转移过来，取二者最大值。
		// 为啥是 减 nums[i]呢，
		dp1[i+1] = max(dp1[i], dp2[i]-nums[i])
		dp2[i+1] = max(dp2[i], dp1[i]+nums[i])
	}
	return int64(max(dp2[n], dp1[n]))
}
