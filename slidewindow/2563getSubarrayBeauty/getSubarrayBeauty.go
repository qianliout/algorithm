package main

import (
	"fmt"
)

func main() {
	fmt.Println(getSubarrayBeauty([]int{1, -1, -3, -2, 3}, 3, 2))
	fmt.Println(getSubarrayBeauty([]int{-3, 1, 2, -3, 0, -3}, 2, 1))
	fmt.Println(getSubarrayBeauty1([]int{-3, 1, 2, -3, 0, -3}, 2, 1))
}

// 值域很小，可以用计数排序的思想
func getSubarrayBeauty1(nums []int, k int, x int) []int {
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
				ans[i] = min(j-base, 0) // 只要负数
				break
			}
		}

		cnt[nums[i]+base]-- // 离开窗口
	}
	return ans
}

// -50 <= nums[i] <= 50
func getSubarrayBeauty(nums []int, k int, x int) []int {
	base := 50 // 防止负数下标
	win := make([]int, 101)
	le, ri, n := 0, 0, len(nums)
	ans := make([]int, 0)
	for le <= ri && ri < n {
		win[nums[ri]+base]++
		ri++
		if ri < k {
			continue
		}
		// 更新答案
		find := x
		for j, c := range win {
			find -= c
			if find <= 0 {
				ans = append(ans, min(0, j-base)) // 只要负数
				break
			}
		}
		// 出窗口了
		for ri-le >= k {
			win[nums[le]+base]--
			le++
		}
	}

	return ans
}
