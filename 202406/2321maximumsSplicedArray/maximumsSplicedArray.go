package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumsSplicedArray([]int{20, 40, 20, 70, 30}, []int{50, 20, 50, 40, 20}))
}

func maximumsSplicedArray(nums1 []int, nums2 []int) int {
	return max(help(nums1, nums2), help(nums2, nums1))
}

func help(nums1 []int, nums2 []int) int {
	n := len(nums1)

	sum := 0
	for _, ch := range nums1 {
		sum += ch
	}

	diff := make([]int, n)
	for i := 0; i < n; i++ {
		diff[i] = nums2[i] - nums1[i]
	}

	return sum + maxDp(diff)
}

// 最大子数组和
func maxDp(nums []int) int {
	n := len(nums)
	mx := make([]int, n)
	mx[0] = nums[0]
	ans := mx[0]
	for i := 1; i < n; i++ {
		mx[i] = max(mx[i-1]+nums[i], nums[i])
		ans = max(ans, mx[i])
	}
	return ans
}
