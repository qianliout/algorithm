package main

import (
	"fmt"
)

func main() {
	// fmt.Println(dominantIndex([]int{1, 2, 3, 4}))
	fmt.Println(dominantIndex([]int{3, 6, 1, 0}))
}

// 请你找出数组中的最大元素并检查它是否 至少是数组中每个其他数字的两倍 。如果是，则返回 最大元素的下标 ，否则返回 -1 。

func dominantIndex(nums []int) int {
	if len(nums) < 2 {
		return -1
	}

	a, b := 0, -1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[a] {
			a, b = i, a
		} else if b == -1 || nums[i] > nums[b] {
			b = i
		}
	}

	if nums[a] >= nums[b]*2 {
		return a
	}
	return -1
}
