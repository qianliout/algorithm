package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxResult([]int{1, -1, -2, 4, -7, 3}, 2))
	fmt.Println(maxResult([]int{1, -5, -20, 4, -1, 3, -6, -3}, 2))
}

// 超时
func maxResult1(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)

	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + nums[i]
		for j := i - 1; j >= 0 && j >= i-k; j-- {
			dp[i] = max(dp[i], dp[j]+nums[i])
		}
	}
	return dp[n-1]
}

func maxResult(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	st := make([]int, 0)
	for i := 0; i < n; i++ {
		// 不在窗口内了
		for len(st) > 0 && st[0]+k < i {
			st = st[1:]
		}
		if len(st) > 0 {
			dp[i] = dp[st[0]] + nums[i]
		}
		for len(st) > 0 && dp[st[len(st)-1]] < dp[i] {
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return dp[n-1]
}
