package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 4, 3, 5, 4, 7, 2}
	fmt.Printf("Input: %v\n", nums)
	fmt.Printf("LIS length: %d\n", lengthOfLIS(nums))
	fmt.Printf("Number of LIS: %d\n", findNumberOfLIS(nums))
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	f := make([]int, n+1)
	ans := 0
	for i := 0; i < n; i++ {
		f[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				f[i] = max(f[i], f[j]+1)
			}
		}
		ans = max(ans, f[i])
	}
	return ans
}

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	// 表示以nums[i]结尾的最长递增子序列的长度
	f := make([]int, n+1)

	// 表示最长递增子序列的长度等 f[i]时的个数
	cnt := make([]int, n+1)
	mx := 0
	for i := 0; i < n; i++ {
		f[i] = 1
		cnt[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				if f[j]+1 > f[i] {
					f[i] = f[j] + 1
					cnt[i] = cnt[j]
				} else if f[j]+1 == f[i] {
					cnt[i] += cnt[j]
				}
			}
		}
		mx = max(mx, f[i])
	}
	ans := 0
	for i := 0; i < n; i++ {
		if f[i] == mx {
			ans += cnt[i]
		}
	}
	return ans
}

/*
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
*/
