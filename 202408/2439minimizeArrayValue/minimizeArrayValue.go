package main

import (
	"slices"
)

func main() {

}

func minimizeArrayValue(nums []int) int {
	mi, mx := slices.Min(nums), slices.Max(nums)
	le, ri := mi, mx+1
	for le < ri {
		mid := le + (ri-le)/2
		if mid >= mi && mid <= mx && check(nums, mid) {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	return le
}

func check(nums []int, limit int) bool {
	extra := 0
	for i := len(nums) - 1; i > 0; i-- {
		extra = max(nums[i]+extra-limit, 0)
	}
	return nums[0]+extra <= limit
}

func check2(nums []int, limit int) bool {
	extra := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if extra+nums[i] > limit {
			extra = nums[i] + extra - limit
		} else {
			extra = 0
		}
	}
	return extra == 0
}
