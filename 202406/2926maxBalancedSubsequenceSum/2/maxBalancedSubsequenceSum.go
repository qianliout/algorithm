package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxBalancedSubsequenceSum([]int{3, 3, 5, 6}))
	fmt.Println(maxBalancedSubsequenceSum([]int{5, -1, -3, 8}))
	fmt.Println(maxBalancedSubsequenceSum([]int{-2, -1}))
}

/*
nums[i]-nums[j] >=i-j
nums[i]-i >= nums[j]-j
把nums[i]-i 定义成一个整体 b
*/
// 不用数状数组的话会超时，结果是对的
func maxBalancedSubsequenceSum(nums []int) int64 {
	n := len(nums)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = nums[i] - i
	}
	ans := math.MinInt
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = nums[i]
		for j := i - 1; j >= 0; j-- {
			if b[j] <= b[i] {
				dp[i] = max(dp[i], dp[j]+nums[i])
			}
		}
		ans = max(ans, dp[i])
	}

	return int64(ans)
}

// 树状数组模板（维护前缀最大值）
type fenwick []int

func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] = max(f[i], val)
	}
}

func (f fenwick) preMax(i int) int {
	mx := math.MinInt
	for ; i > 0; i &= i - 1 {
		mx = max(mx, f[i])
	}
	return mx
}
