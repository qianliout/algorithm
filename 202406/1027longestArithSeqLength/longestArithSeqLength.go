package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestArithSeqLength([]int{20, 1, 15, 3, 10, 5, 8}))
	fmt.Println(longestArithSeqLength([]int{9, 4, 7, 2, 10}))
	fmt.Println(longestArithSeqLength([]int{22, 8, 57, 41, 36, 46, 42, 28, 42, 14, 9, 43, 27, 51, 0, 0, 38, 50, 31, 60, 29, 31, 20, 23, 37, 53, 27, 1, 47, 42, 28, 31, 10, 35, 39, 12, 15, 6, 35, 31, 45, 21, 30, 19, 5, 5, 4, 18, 38, 51, 10, 7, 20, 38, 28, 53, 15, 55, 60, 56, 43, 48, 34, 53, 54, 55, 14, 9, 56, 52}))
}

func longestArithSeqLength(nums []int) int {
	n := len(nums)
	dp := make([]map[int]int, n)
	for i := range dp {
		dp[i] = make(map[int]int)
	}
	ans := 0
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			d := nums[i] - nums[j]
			dp[i][d] = max(dp[i][d], dp[j][d]+1)
			ans = max(ans, dp[i][d])
		}
	}
	return ans + 1
}
