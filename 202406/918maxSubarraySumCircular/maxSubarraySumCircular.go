package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubarraySumCircular([]int{1, -2, 3, -2}))
	fmt.Println(maxSubarraySumCircular([]int{3, -2, 2, -3}))
	fmt.Println(maxSubarraySumCircular([]int{5, -3, 5}))
}

// 会超时
func maxSubarraySumCircular1(nums []int) int {
	n := len(nums)

	ans := math.MinInt
	for start := 0; start < n; start++ {
		mx := make([]int, n*2)
		mx[start] = nums[start]
		ans = max(ans, mx[start])
		for i := start + 1; i < start+n; i++ {
			idx := i % n
			mx[i] = max(nums[idx], mx[i-1]+nums[idx])
			ans = max(ans, mx[i])
		}
	}
	return ans
}

// 参考灵神的题解
// 分两种情况，如果没有跨边界，那么就是 maxS
// 如果跨了边界，就是所有元素和减去数组最小的子数组和
// 特别的情况是：如果最小的子数组是整个数据，那么最大的数组就是空，这是不可以的，这种情况下就返回不跨边界的最大子数组就好
func maxSubarraySumCircular(nums []int) int {
	maxS := math.MinInt // 最大子数组和，不能为空
	minS := 0           // 最小子数组和，可以为空
	sum := 0
	mxF, miF := 0, 0
	for _, x := range nums {
		mxF = max(0, mxF) + x
		miF = min(0, miF) + x
		sum += x
		maxS = max(maxS, mxF)
		minS = min(minS, miF)
	}
	if sum == minS {
		return maxS
	}
	return max(maxS, sum-minS)
}
