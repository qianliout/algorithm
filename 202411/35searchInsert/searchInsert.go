package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
}

func searchInsert1(nums []int, target int) int {
	j := sort.SearchInts(nums, target)
	return j
}

func searchInsert(nums []int, target int) int {
	n := len(nums)
	le, ri := 0, n
	for le < ri {
		// 找>=target的左端点
		mid := le + (ri-le)/2
		if mid < n && nums[mid] >= target {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	return le
}
