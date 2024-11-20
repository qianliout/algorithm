package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubarraySumCircular([]int{5, -3, 5}))
}

// 这样写会超时
func maxSubarraySumCircular1(nums []int) int {
	ans := math.MinInt64
	n := len(nums)
	for i := 0; i < n; i++ {
		ans = max(ans, maxSubArray(nums, i))
	}
	return ans
}

func maxSubArray(nums []int, start int) int {
	n := len(nums)
	dp := make([]int, n*2)
	dp[start] = nums[start]
	ans := nums[start]
	for i := start + 1; i < n+start; i++ {
		idx := i % n
		dp[i] = max(dp[i-1]+nums[idx], nums[idx])
		ans = max(ans, dp[i])
	}
	return ans
}

// 参考灵神的题解
// 分两种情况，如果没有跨边界，那么就是 maxS
// 如果跨了边界，就是所有元素和减去数组最小的子数组和
// 特别的情况是：如果最小的子数组是整个数据，那么最大的数组就是空，这是不可以的，
// 这种情况下就返回不跨边界的最大子数组就好
func maxSubarraySumCircular(nums []int) int {
	maxS := math.MinInt   // 最大子数组和，不能为空
	minS := math.MaxInt64 // 最小子数组和，可以为空
	sum := 0              // 元素总和
	mxF, miF := 0, 0      // db数组，也就是之前的最大子数组和，最小子数组和
	for _, x := range nums {
		mxF = max(mxF+x, x)
		miF = min(miF+x, x)
		sum += x
		maxS = max(maxS, mxF)
		minS = min(minS, miF)
	}
	if sum == minS {
		return maxS
	}
	return max(maxS, sum-minS)
}
