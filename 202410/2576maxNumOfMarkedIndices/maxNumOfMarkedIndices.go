package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(maxNumOfMarkedIndices([]int{9, 2, 5, 4}))
	fmt.Println(maxNumOfMarkedIndices([]int{42, 83, 48, 10, 24, 55, 9, 100, 10, 17, 17, 99, 51, 32, 16, 98, 99, 31, 28, 68, 71, 14, 64, 29, 15, 40}))
}

func maxNumOfMarkedIndices2(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	i := 0
	for _, ch := range nums[(n+1)>>1:] {
		if nums[i]*2 <= ch {
			i++
		}
	}
	return i * 2
}

// 二分
func maxNumOfMarkedIndices(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	left, right := 0, n
	for left < right {
		mid := left + (right-left+1)>>1
		if mid >= 0 && mid < n && mid <= n/2 && check(nums, mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return left * 2
}

func check(nums []int, k int) bool {
	n := len(nums)
	for i := 0; i < k; i++ {
		if nums[i]*2 > nums[n-k+i] {
			return false
		}
	}
	return true
}
