package main

import (
	"sort"
)

func main() {

}

func minOperations(nums []int, target int) int {
	all := 0
	for _, ch := range nums {
		all += ch
	}
	if all < target {
		return -1
	}
	if all == target {
		return 0
	}
	sort.Ints(nums)
	ans := 0
	for target > 0 && len(nums) > 0 {
		// 这里的排序可有可无
		sort.Ints(nums)
		// 把数组里最后最大的数字拿出来
		b := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		// 如果不用这个数字，靠前面的数字都能凑出来target，这个数字直接扔掉
		if all-b >= target {
			all -= b
		} else if b > target { // 如果这个数字本身就比target大
			ans++
			nums = append(nums, b/2, b/2) // 拆完了塞回去数组尾部
			// 如果这个数字刚好是target或者小于target
		} else {
			target -= b
			all -= b
		}
	}
	return ans
}
