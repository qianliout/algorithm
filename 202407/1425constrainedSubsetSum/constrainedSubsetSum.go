package main

import (
	"fmt"
)

func main() {
	fmt.Println(constrainedSubsetSum([]int{10, 2, -10, 5, 20}, 2))
}

func constrainedSubsetSum(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = nums[i]
	}
	ans := nums[0]
	st := make([]int, 0) // 很多题解都加了哨兵，但是我对这个应用的不熟悉
	for i := 0; i < n; i++ {
		for len(st) > 0 && i-st[0] > k {
			st = st[1:]
		}
		if len(st) > 0 {
			dp[i] = max(0, dp[st[0]]) + nums[i]
		}

		ans = max(ans, dp[i])
		for len(st) > 0 && dp[st[len(st)-1]] <= dp[i] {
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return ans
}
