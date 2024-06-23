package main

import (
	"fmt"
)

func main() {
	fmt.Println(numberOfSubarrays([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2))
}

func numberOfSubarrays(nums []int, k int) int {
	return atMost(nums, k) - atMost(nums, k-1)
}

// 不大于 goal 的非空子数组
func atMost(nums []int, goal int) int {
	// nums里都是正数
	if goal < 0 {
		return 0
	}
	le, ri := 0, 0
	ans := 0
	sum := 0
	for le <= ri && ri < len(nums) {
		sum += nums[ri] & 1
		ri++

		for le <= ri && sum > goal {
			sum -= nums[le] & 1
			le++
		}
		ans += ri - le
	}

	return ans
}
