package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println()
}

/*
给你一个整数数组 nums 和一个整数 k ，找出 nums 中和至少为 k 的 最短非空子数组 ，并返回该子数组的长度。如果不存在这样的 子数组 ，返回 -1 。
子数组 是数组中 连续 的一部分。
*/
func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	sum := make([]int, n+1)
	ans := math.MaxInt / 2
	for i, ch := range nums {
		sum[i+1] = sum[i] + ch
	}
	stark := make([]int, 0)
	// 单点递增的
	// 这里的i < n+1,是一个容器出错的点
	for i := 0; i < len(sum); i++ {
		// 对于当前i 来说，如果前面有一个 j,sum[i]-sum[j]>=k,说明i和j 已经可以组成一个符合要求的子数组，又因为，这个队列是单调递增的，j后面的数所能生成的答案会更短
		// 所以可以计算 j 所能生成的答案，并弹出去，注意这里是弹左边
		for len(stark) > 0 && sum[i]-sum[stark[0]] >= k {
			ans = min(ans, i-stark[0])
			stark = stark[1:]
		}
		// 对于当前的 i 来说，如果前面的 sum[j]大于当前的 i，那么对于后面的数来说，如果能和 j 组成子数组，那么一定能和 i组成更小的子数组，弹右边
		for len(stark) > 0 && sum[stark[len(stark)-1]] >= sum[i] {
			stark = stark[:len(stark)-1]
		}
		stark = append(stark, i)
	}

	if ans >= math.MaxInt/2 {
		return -1
	}
	return ans
}
