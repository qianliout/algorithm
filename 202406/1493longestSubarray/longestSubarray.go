package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 1, 1}))
}

// 必需删除一个元素,优先删除0
func longestSubarray(nums []int) int {
	ans, n := 0, len(nums)
	le, ri, cnt := 0, 0, 0
	for le <= ri && ri < n {
		cnt += nums[ri]
		for ri-le+1 > cnt+1 { // 最多只删一个0
			cnt -= nums[le]
			le++
		}
		ri++
		ans = max(ans, cnt)
	}
	return min(ans, n-1)
}
