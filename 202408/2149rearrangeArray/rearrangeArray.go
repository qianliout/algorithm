package main

import (
	"fmt"
)

func main() {
	fmt.Println(rearrangeArray([]int{3, 1, -2, -5, 2, -4}))
}

func rearrangeArray(nums []int) []int {
	n := len(nums)
	nums1, nums2 := make([]int, 0), make([]int, 0)
	for _, ch := range nums {
		if ch > 0 {
			nums1 = append(nums, ch)
		} else {
			nums2 = append(nums2, ch)
		}
	}
	ans := make([]int, 0)
	for i := 0; i < n/2; i++ {
		ans = append(ans, nums1[i])
		ans = append(ans, nums2[i])
	}
	return ans
}

// 由数目 相等 的正整数和负整数组成。
