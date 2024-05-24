package main

import (
	"fmt"
)

func main() {
	fmt.Println(goodIndices([]int{2, 1, 1, 1, 3, 4, 1}, 2))
	fmt.Println(goodIndices([]int{253747, 459932, 263592, 354832, 60715, 408350, 959296}, 2))
}

func goodIndices(nums []int, k int) []int {

	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	for i := 0; i < n; i++ {
		left[i], right[i] = 1, 1
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] <= nums[i+1] {
			right[i] = right[i+1] + 1
		}
	}
	ans := make([]int, 0)
	for i := k; i+k < n; i++ {
		if left[i-1] >= k && right[i+1] >= k {
			ans = append(ans, i)
		}
	}

	return ans
}
