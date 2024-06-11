package main

import (
	"fmt"
)

func main() {
	fmt.Println(48 & 10)
	fmt.Println(longestNiceSubarray([]int{1, 3, 8, 48, 10}))
}

func longestNiceSubarray(nums []int) int {
	ans, n := 1, len(nums)
	le, ri := 0, 0
	set := 0
	for le <= ri && ri < n {
		c := nums[ri]
		for set&c > 0 { // 有交集
			set ^= nums[le] // 从 set 中去掉集合 nums[left]
			le++
		}
		set = set | c
		ri++
		ans = max(ans, ri-le)
	}
	return ans
}
