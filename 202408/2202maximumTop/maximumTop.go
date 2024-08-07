package main

import (
	"slices"
)

func main() {

}

func maximumTop(nums []int, k int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		if k&1 == 0 {
			return nums[0]
		}
		return -1
	}
	if k <= 1 {
		return nums[k]
	}
	if k > n {
		return slices.Max(nums)
	}
	if k == n {
		return slices.Max(nums[:n-1])
	}
	return max(slices.Max(nums[:k-1]), nums[k])
}
