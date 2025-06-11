package main

import (
	"fmt"
)

func main() {
	fmt.Println(rob([]int{1}))
}

func rob(nums []int) int {
	// 这里的判断容易出错

	if len(nums) == 1 {
		return nums[0]
	}
	// return max(dfs(nums[1:]), dfs(nums[:len(nums)-1]))
	return max(dp(nums[1:]), dp(nums[:len(nums)-1]))
}

// 直接这样写会超时
func dfs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	y := nums[0]
	if len(nums) >= 2 {
		y += dfs(nums[2:])
	}
	n := dfs(nums[1:])
	return max(y, n)
}

func dp(nums []int) int {
	n := len(nums)
	dp1 := make([]int, n+1) // yes
	dp2 := make([]int, n+1) // no
	dp1[0] = nums[0]
	for i := 1; i < n; i++ {
		dp1[i] = max(dp1[i], dp2[i-1]+nums[i])
		dp2[i] = max(dp2[i], dp1[i-1], dp2[i-1])
	}
	return max(dp1[n-1], dp2[n-1])
}
