package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestArithSeqLength([]int{3, 6, 9, 12}))
}

func longestArithSeqLength(nums []int) int {
	n := len(nums)
	f := make([]map[int]int, n)
	mx := 0
	for i := 0; i < n; i++ {
		if f[i] == nil {
			f[i] = make(map[int]int)
		}

		for j := i - 1; j >= 0; j-- {
			sub := nums[i] - nums[j]
			f[i][sub] = max(f[i][sub], f[j][sub]+1)
			mx = max(mx, f[i][sub])
		}
	}
	return mx + 1
}

// 给你一个整数数组 nums，返回 nums 中最长等差子序列的长度
