package main

import (
	"slices"
)

func main() {

}

// 结果正确，复杂度高
func kIncreasing(arr []int, k int) int {
	nums := make([][]int, k)
	for i, ch := range arr {
		nums[i%k] = append(nums[i%k], ch)
	}
	ans := 0
	for _, num := range nums {
		ans += len(num) - help(num)
	}
	return ans
}

// 最大递增子序列
func help(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] <= nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	return slices.Max(dp)
}
