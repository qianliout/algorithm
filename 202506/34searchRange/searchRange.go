package main

import (
	"fmt"
)

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}

func searchRange(nums []int, target int) []int {
	n := len(nums)
	left, right := 0, n
	// 左边
	ans := make([]int, 2)
	ans[0], ans[1] = -1, -1

	for left < right {
		mid := left + (right-left)/2
		//
		if mid < n && nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if left >= 0 && left < n && nums[left] == target {
		ans[0] = left
	}
	left, right = 0, n
	for left < right {
		mid := left + (right-left+1)/2
		if mid < n && nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	if left >= 0 && left < n && nums[left] == target {
		ans[1] = left
	}
	return ans
}
