package main

import (
	"math"
)

func main() {

}

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	ans := math.MinInt64
	for i := 0; i < n; i++ {
		if i == 0 {
			ans = max(ans, nums[i])
		} else {
			dp[i] = max(dp[i-1]+nums[i], nums[i])
			ans = max(ans, dp[i])
		}
	}
	return ans
}
