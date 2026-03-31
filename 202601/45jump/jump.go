package main

import (
	"math"
)

func main() {}

func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	inf := math.MaxInt64 / 100
	for i := 0; i < n; i++ {
		dp[i] = inf
	}
	dp[0] = 0
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] != inf && i-j <= nums[j] {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	if dp[n-1] == inf {
		return -1
	}
	return dp[n-1]
}
