package main

import (
	"sort"
)

func main() {

}

func makeSimilar(nums []int, target []int) int64 {
	f(nums)
	f(target)
	ans := 0
	n := len(nums)
	// for i := 0; i < n; i++ {
	// 	ans += abs(nums[i] - target[i])
	// }
	// return int64(ans) / 4
	// 题目中数据都是正数
	for i := 0; i < n; i++ {
		if nums[i] < 0 {
			ans += abs(-nums[i] - (-target[i]))
		} else {
			ans += abs(nums[i] - target[i])
		}
	}
	// 令 nums[i] = nums[i] + 2 且
	// 令 nums[j] = nums[j] - 2 。
	// 这两步操作一起才算一步，
	return int64(ans) / 4
}

func f(a []int) {
	for i, ch := range a {
		if ch&1 == 1 {
			a[i] = -ch
		}
	}

	sort.Ints(a)

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
