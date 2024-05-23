package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxWidthRamp([]int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}))
	fmt.Println(maxWidthRamp([]int{6, 0, 8, 2, 1, 5}))
}

func maxWidthRamp(nums []int) int {
	stark := make([]int, 0)
	// 以nums[0]为栈顶元素，维护一个单调栈
	for i := range nums {
		if len(stark) == 0 || nums[stark[len(stark)-1]] > nums[i] {
			stark = append(stark, i)
		}
	}
	ans := 0
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stark) > 0 && nums[i] >= nums[stark[len(stark)-1]] {
			ans = max(ans, i-stark[len(stark)-1])
			stark = stark[:len(stark)-1]
		}
	}
	return ans
}
