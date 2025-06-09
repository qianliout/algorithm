package main

import (
	"fmt"
)

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
}

func firstMissingPositive(nums []int) int {
	n := len(nums)
	// 这是这一个题目不容易想到的地方
	nums = append(nums, 0)
	for i := range nums {
		if nums[i] < 0 || nums[i] > n {
			nums[i] = 0
		}
	}
	le := 0
	for le <= n {
		for nums[le] != nums[nums[le]] && nums[le] != le {
			nums[le], nums[nums[le]] = nums[nums[le]], nums[le]
		}
		le++
	}
	for i := 1; i <= n; i++ {
		if nums[i] != i {
			return i
		}
	}
	return n + 1
}
