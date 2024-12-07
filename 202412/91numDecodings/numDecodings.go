package main

import (
	"fmt"
)

func main() {
	fmt.Println(numDecodings1("12"))
}

func numDecodings1(s string) int {
	n := len(s)
	nums := make([]int, n)
	for i := range s {
		nums[i] = int(s[i]) - int('0')
	}
	if n == 0 || nums[0] == 0 {
		return 0
	}
	f := make([]int, n+5)
	f[0], f[1] = 1, 1
	for i := 1; i < n; i++ {
		if nums[i] == 0 {
			if nums[i-1] == 0 || nums[i-1] > 2 {
				return 0
			}
			f[i+1] = f[i-1]
			continue
		}
		if nums[i-1] == 1 || (nums[i-1] == 2 && nums[i] <= 6 && nums[i] >= 0) {
			f[i+1] = f[i-1] + f[i]
			continue
		}
		f[i+1] = f[i]
	}
	return f[n]
}

func numDecodings(s string) int {
	nums := make([]int, 0)
	for i := range s {
		nums = append(nums, int(s[i])-48)
	}
	// 边界条件
	if len(nums) == 0 || nums[0] == 0 {
		return 0
	}

	// dp[i] 表示 nums[:i)的解法
	dp := make([]int, len(nums)+1)
	dp[0], dp[1] = 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			if nums[i-1] == 1 || nums[i-1] == 2 {
				dp[i+1] = dp[i-1]
				continue
			}
			return 0
		}

		if nums[i-1] == 1 || (nums[i-1] == 2 && nums[i] <= 6 && nums[i] >= 0) {
			dp[i+1] = dp[i-1] + dp[i]
			continue
		}

		dp[i+1] = dp[i]
	}
	return dp[len(nums)]
}
