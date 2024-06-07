package main

import (
	"fmt"
)

func main() {
	fmt.Println(minSumOfLengths([]int{3, 2, 2, 4, 3}, 3))
	fmt.Println(minSumOfLengths([]int{7, 3, 4, 7}, 7))
	fmt.Println(minSumOfLengths([]int{4, 3, 2, 6, 2, 3, 4}, 6))
	fmt.Println(minSumOfLengths([]int{5, 5, 4, 4, 5}, 3))
	fmt.Println(minSumOfLengths([]int{1, 1, 1, 2, 2, 2, 4, 4}, 6))
}

func minSumOfLengths(nums []int, target int) int {
	n := len(nums)

	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = n + 1
	}
	ans := n + 1
	le, ri := 0, 0
	sum := 0
	for le <= ri && ri < n {
		sum += nums[ri]
		ri++
		for sum > target {
			sum -= nums[le]
			le++
		}
		if sum == target {
			ans = min(ans, ri-le+dp[le])
			dp[ri] = min(dp[ri-1], ri-le)
		} else {
			dp[ri] = dp[ri-1]
		}
	}
	if ans == n+1 {
		return -1
	}
	return ans
}
