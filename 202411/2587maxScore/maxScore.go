package main

import (
	"sort"
)

func main() {

}

func maxScore(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}
	for i, ch := range nums {
		if ch <= 0 {
			return i
		}
	}
	return 0
}
