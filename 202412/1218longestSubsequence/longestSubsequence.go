package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestSubsequence([]int{1, 2, 3, 4}, 1))
}

func longestSubsequence(arr []int, difference int) int {
	mx := 0
	f := make(map[int]int)
	for _, x := range arr {
		pre := x - difference
		f[x] = max(1, f[pre]+1)
		mx = max(mx, f[x])
	}
	return mx
}

// 给你一个整数数组 arr 和一个整数 difference，请你找出并返回 arr 中最长等差子序列的长度，该子序列中相邻元素之间的差等于 difference 。
