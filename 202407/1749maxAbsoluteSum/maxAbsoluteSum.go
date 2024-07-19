package main

import (
	"slices"
)

func main() {

}

func maxAbsoluteSum(nums []int) int {
	n := len(nums)
	mx, mi := make([]int, n), make([]int, n)
	mx[0], mi[0] = nums[0], nums[0]
	for i, ch := range nums {
		if i == 0 {
			continue
		}
		mx[i] = max(ch, mx[i-1]+ch)
		mi[i] = min(ch, mi[i-1]+ch)
	}
	return max(0, abs(slices.Max(mx)), abs(slices.Min(mi)))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
