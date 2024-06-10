package main

import (
	"fmt"
)

func main() {
	fmt.Println(getSubarrayBeauty([]int{1, -1, -3, -2, 3}, 3, 2))
}

// 值域很小，可以用计数排序的思想
func getSubarrayBeauty(nums []int, k int, x int) []int {
	base := 50
	cnt := make([]int, 101)
	for _, ch := range nums[:k-1] {
		cnt[ch+base]++
	}
	n := len(nums)
	ans := make([]int, n-k+1)
	after := nums[k-1:]
	for i, ch := range after {
		cnt[ch+base]++
		left := x
		// 只找负数
		for j, c := range cnt[:base] {
			left -= c
			if left <= 0 {
				ans[i] = j - base
				break
			}
		}

		cnt[nums[i]+base]-- // 离开窗口
	}
	return ans
}
