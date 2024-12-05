package main

import (
	"slices"
)

func main() {

}

func deleteAndEarn(nums []int) int {
	mx := slices.Max(nums)
	// 技巧：开一个较大的数组，这样就不用对 0，1，2，单独判断了
	f := make([]int, mx+10)
	for _, ch := range nums {
		f[ch]++
	}
	// 表示前 i 个数，能够获取的最大值
	// 技巧：开一个较大的数组，这样就不用对 0，1，2，单独判断了
	dp := make([]int, mx+10) // yes
	dp[1], dp[2] = f[1]*1, f[2]*2

	for i := 2; i <= mx; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+f[i]*i)
	}
	return dp[mx]
}
