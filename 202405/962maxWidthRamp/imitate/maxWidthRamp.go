package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxWidthRamp([]int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}))
	fmt.Println(maxWidthRamp([]int{6, 0, 8, 2, 1, 5}))
}

func maxWidthRamp(nums []int) int {
	mem := make(map[int]int)
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans = max(ans, imitate(nums, i, mem))
	}
	return ans
}

func imitate(nums []int, start int, mem map[int]int) int {
	if start >= len(nums) {
		return 0
	}
	if va, ok := mem[start]; ok {
		return va
	}

	ans := 0
	for i := len(nums) - 1; i > start; i-- {
		if nums[i] >= nums[start] {
			ans = max(ans, i-start)
			break
		}
	}
	mem[start] = ans
	return ans
}
