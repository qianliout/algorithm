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

// dp的写法
// https://leetcode.cn/problems/house-robber/submissions/
func check2(nums []int, mid int, k int) bool {
	n := len(nums)
	// 定义 f[i] 表示从 nums[0] 到 nums[i] 中偷金额不超过 mx 的房屋，最多能偷多少间房屋。如果 f[n−1]≥k 则表示答案至多为 mx，否则表示答案必须超过 mx。
	f := make([]int, n+2)
	// 初值
	for i := 0; i < n; i++ {
		ch := nums[i]
		if ch > mid {
			f[i+2] = f[i+1]
		} else {
			f[i+2] = max(f[i+1], f[i]+1)
		}
	}
	return f[n+1] >= k
}

func check(nums []int, mid int, k int) bool {
	n := len(nums)
	// 定义 f[i] 表示从 nums[0] 到 nums[i] 中偷金额不超过 mx 的房屋，最多能偷多少间房屋。如果 f[n−1]≥k 则表示答案至多为 mx，否则表示答案必须超过 mx。
	// f := make([]int, n+2)
	f0, f1 := 0, 0
	// 初值
	for i := 0; i < n; i++ {
		ch := nums[i]
		if ch > mid {
			f0 = f1
		} else {
			f0, f1 = f1, max(f1, f0+1)
		}
	}
	return f1 >= k
}
