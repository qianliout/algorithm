package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 4, 3, 5, 4, 7, 2}
	fmt.Printf("Input: %v\n", nums)
	fmt.Printf("LIS length: %d\n", lengthOfLIS(nums))
	fmt.Printf("Number of LIS: %d\n", findNumberOfLIS(nums))

	// 调试信息
	debugFindNumberOfLIS(nums)
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

func debugFindNumberOfLIS(nums []int) {
	n := len(nums)
	if n == 0 {
		return
	}

	lengths := make([]int, n)
	counts := make([]int, n)

	for i := 0; i < n; i++ {
		lengths[i] = 1
		counts[i] = 1
	}

	maxLen := 1

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if lengths[j]+1 > lengths[i] {
					lengths[i] = lengths[j] + 1
					counts[i] = counts[j]
				} else if lengths[j]+1 == lengths[i] {
					counts[i] += counts[j]
				}
			}
		}
		maxLen = max(maxLen, lengths[i])
	}

	fmt.Printf("lengths: %v\n", lengths)
	fmt.Printf("counts:  %v\n", counts)
	fmt.Printf("maxLen:  %d\n", maxLen)

	// 打印以每个位置结尾的LIS信息
	for i := 0; i < n; i++ {
		if lengths[i] == maxLen {
			fmt.Printf("Position %d (value=%d): length=%d, count=%d\n", i, nums[i], lengths[i], counts[i])
		}
	}
}
