package main

import (
	"fmt"
)

func main() {
	fmt.Println(canSplitArray([]int{2, 1, 3}, 5))
	fmt.Println(canSplitArray([]int{1, 1}, 3))
}

func canSplitArray(nums []int, m int) bool {
	if len(nums) <= 2 {
		return true
	}
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i]+nums[i-1] >= m {
			return true
		}
	}
	return false
}
