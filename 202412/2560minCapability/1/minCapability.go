package main

import (
	"slices"
)

func main() {

}

func minCapability(nums []int, k int) int {
	mx := slices.Max(nums)
	le, ri := 0, mx+1
	for le < ri {
		// 最小值，查左端点
		mid := le + (ri-le)/2
		if mid >= 0 && mid < mx+1 && check(nums, mid, k) {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	// 检测一下
	if check(nums, le, k) {
		return le
	}
	return -1
}

// 贪心的写法，能偷就尽快偷 但是到底对不对，没有想明白
func check(nums []int, mid int, k int) bool {
	cnt := 0
	i := 0
	for i < len(nums) {
		if nums[i] <= mid {
			cnt++
			i += 2
		} else {
			i++
		}
	}
	return cnt >= k
}
