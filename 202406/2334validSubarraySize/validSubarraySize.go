package main

import (
	"fmt"
)

func main() {
	// fmt.Println(validSubarraySize([]int{1, 3, 4, 3, 1}, 6))
	fmt.Println(validSubarraySize([]int{6, 5, 6, 5, 8}, 7))
}

func validSubarraySize(nums []int, threshold int) int {
	n := len(nums)
	// 求nums[i]做为最小值，能向左走多远
	left := make([]int, n)
	for i := range left {
		left[i] = -1 // 默认能走到最左侧，就是-1
	}
	st := make([]int, 0) // 单调栈

	for i, ch := range nums {
		for len(st) > 0 && ch <= nums[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			left[i] = st[len(st)-1]
		}
		st = append(st, i)
	}

	right := make([]int, n)
	for i := range right {
		right[i] = n
	}
	st = st[:0]

	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && nums[i] <= nums[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			right[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	for i, ch := range nums {
		k := right[i] - left[i] - 1
		if ch > threshold/k {
			return k
		}
	}
	return -1

}
