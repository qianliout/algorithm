package main

import (
	"fmt"
)

func main() {
	fmt.Println(lenLongestFibSubseq([]int{1, 2, 3, 4, 5, 6, 7, 8}))  // 5
	fmt.Println(lenLongestFibSubseq([]int{1, 3, 7, 11, 12, 14, 18})) // 8
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
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dp[i][j] = 2
		}
	}

	ans := 0
	// 到索引的映射
	// 题目中已经明确：给定一个严格递增的正整数数组形成序列 arr ，找到 arr 中最长的斐波那契式的子序列的长度。如果一个不存在，返回  0 。
	intMap := make(map[int]int)

	for i := 0; i < n; i++ {
		intMap[nums[i]] = i
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			diff := nums[j] - nums[i]
			// 因为是严格递增的，说明没有重复值
			// a+b=c
			if va, ok := intMap[diff]; ok && va < i {
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
