package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxStrength([]int{8, 8, -6, -4, -6, -4, -6, -5, 0, -9, -6, -3}))
}

func maxStrength(nums []int) int64 {
	n := len(nums)
	mi, mx := nums[0], nums[0]

	for i := 1; i < n; i++ {
		tm := mx
		mx = max(mx*nums[i], nums[i], mi*nums[i], mx)
		mi = min(tm*nums[i], nums[i], mi*nums[i], mi)
	}
	return int64(mx)
}
