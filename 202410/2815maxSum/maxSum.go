package main

import (
	"slices"
)

func main() {

}

func maxSum(nums []int) int {
	n := len(nums)
	ans := -1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if slices.Max(get(nums[i])) == slices.Max(get(nums[j])) {
				ans = max(ans, nums[i]+nums[j])
			}
		}
	}
	return ans
}

func get(a int) []int {
	ans := make([]int, 0)
	for a > 0 {
		ans = append(ans, a%10)
		a /= 10
	}
	return ans
}
