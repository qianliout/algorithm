package main

import (
	"fmt"
)

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}

func searchRange(nums []int, target int) []int {
	n := len(nums)
	le, ri := 0, n
	for le < ri {
		// 左端点
		mid := le + (ri-le)/2
		if mid >= 0 && mid < n && nums[mid] >= target {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	left := le

	le, ri = 0, n
	for le < ri {
		mid := le + (ri-le+1)/2
		if mid >= 0 && mid < n && nums[mid] <= target {
			le = mid
		} else {
			ri = mid - 1
		}
	}
	right := le
	if left < 0 || right < 0 || left >= n || right >= n ||
		nums[left] != target || nums[right] != target {
		return []int{-1, -1}
	}
	return []int{left, right}
}
