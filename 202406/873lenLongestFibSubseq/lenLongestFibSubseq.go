package main

import (
	"fmt"
)

func main() {
	fmt.Println(lenLongestFibSubseq([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(lenLongestFibSubseq([]int{1, 3, 7, 11, 12, 14, 18}))
}

// 会超时
func lenLongestFibSubseq2(nums []int) int {
	n := len(nums)
	// 表示以A[i],A[j]结尾的斐波那契数列的最大长度
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dp[i][j] = 2
			for k := i - 1; k >= 0; k-- {
				if nums[k]+nums[i] == nums[j] {
					dp[i][j] = max(dp[i][j], dp[k][i]+1)
				}
			}
			ans = max(ans, dp[i][j])
		}
	}
	if ans > 2 {
		return ans
	}
	return 0
}

func lenLongestFibSubseq(nums []int) int {
	n := len(nums)
	// 表示以A[i],A[j]结尾的斐波那契数列的最大长度
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	ans := 0
	diffMap := make(map[int]int)
	for i := 0; i < n; i++ {
		diffMap[nums[i]] = i
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dp[i][j] = 2
			diff := nums[j] - nums[i]
			if va, ok := diffMap[diff]; ok && va < i {
				dp[i][j] = max(dp[i][j], dp[va][i]+1)
			}
			ans = max(ans, dp[i][j])
		}
	}
	if ans > 2 {
		return ans
	}
	return 0
}
