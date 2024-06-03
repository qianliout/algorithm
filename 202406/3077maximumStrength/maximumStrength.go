package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maximumStrength([]int{1, 2, 3, -1, 2}, 3))
}

func maximumStrength2(nums []int, k int) int64 {
	n := len(nums)
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	sum := make([]int, n+1)
	for i, ch := range nums {
		sum[i+1] = ch + sum[i]
	}
	for i := 1; i <= k; i++ {
		// 初值,求最大值，这种情况是不可能的，所以赋一个最小值
		dp[i][i-1] = math.MinInt
		// 这里没有能理解
		w := k - i + 1
		if i%2 == 0 {
			w = -w
		}
		mx := math.MinInt
		for j := i; j <= n-k+i; j++ {
			mx = max(mx, dp[i-1][j-1]-sum[j-1]*w)
			dp[i][j] = max(dp[i][j-1], sum[j]*w+mx)
		}
	}
	return int64(dp[k][n])
}

// 还是有问题,不能正确得到答案
func maximumStrength(nums []int, k int) int64 {
	n := len(nums)
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	sum := make([]int, n+1)
	for i, ch := range nums {
		sum[i+1] = ch + sum[i]
	}
	for i := 1; i <= k; i++ {
		// 初值,求最大值，这种情况是不可能的，所以赋一个最小值
		dp[i][i-1] = math.MinInt
		// 这里没有能理解
		w := k - i + 1
		if i%2 == 0 {
			w = -w
		}
		mx := math.MinInt

		for j := i; j <= n-k+i; j++ {
			mx = max(mx, dp[i-1][j-1]-sum[j-1]*w)
			dp[i][j] = max(dp[i][j-1], sum[j]*w+mx)
		}
	}
	return int64(dp[k][n])
}
