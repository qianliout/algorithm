package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumMountainRemovals([]int{1, 3, 1}))
	fmt.Println(minimumMountainRemovals([]int{100, 92, 89, 77, 74, 66, 64, 66, 64}))
}

func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	pre := make([]int, n)
	suf := make([]int, n)
	for i := 0; i < n; i++ {
		pre[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				pre[i] = max(pre[i], pre[j]+1)
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		suf[i] = 1
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				suf[i] = max(suf[i], suf[j]+1)
			}
		}
	}
	ans := 0
	for i := 1; i < n-1; i++ {
		// 这个判断容易忘记
		if pre[i] > 1 && suf[i] > 1 {
			ans = max(ans, pre[i]+suf[i]-1)
		}
	}
	return n - ans
}
