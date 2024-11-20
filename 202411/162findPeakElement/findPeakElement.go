package main

import (
	"math"
)

func main() {

}

func findPeakElement(nums []int) int {
	nums = append(nums, math.MinInt64)
	n := len(nums)

	for i := 0; i < n; i++ {

		if i == 0 && nums[i] > nums[i+1] {
			return 0
		}
		if i == 0 {
			continue
		}

		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}
	return n - 1
}
