package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

func longestConsecutive2(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	dp := make(map[int]int)
	ans := 0
	for i := 0; i < n; i++ {
		c := nums[i]
		dp[c] = dp[c-1] + 1
		ans = max(ans, dp[c])
	}
	return ans
}

func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	dp := make(map[int]int)
	dp2 := make(map[int]int)
	ans := 0
	for i := 0; i < n; i++ {
		c := nums[i]
		dp[c] = dp[c-1] + 1
	}
	for i := n - 1; i >= 0; i-- {
		c := nums[i]
		dp2[c] = dp2[c+1] + 1
	}
	for _, c := range nums {
		ans = max(ans, dp[c-1]+dp2[c+1]+1)
	}

	return ans
}
