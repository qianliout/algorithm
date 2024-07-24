package main

import (
	"fmt"
)

func main() {
	fmt.Println(canBeIncreasing([]int{100, 21, 100}))
}

func canBeIncreasing(nums []int) bool {
	n := len(nums)
	dp := make([]int, n)
	mx := 1 // 求最长上升子序列
	for i := 0; i < n; i++ {
		dp[i] = 1 // 初值
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		mx = max(mx, dp[i])
	}
	// [1,2,3] 这种情况下，mx==3,但是也是符和要求的,所以要 mx+1>=len(nums)
	return mx+1 >= len(nums)
}
