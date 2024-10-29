package main

import (
	"math"
	"sort"
)

func countWays(ranges [][]int) int {
	// 设合并后有 m 个大区间，那么每个大区间都可以分到第一个组或者第二个组，每个大区间都有 2 个方案。
	// 由于不同的大区间之间互相独立，根据乘法原理，方案数为 2m次方。

	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][0] == ranges[j][0] {
			return ranges[i][1] < ranges[j][1]
		}
		return ranges[i][0] < ranges[j][0]
	})
	n := len(ranges)
	right := -1
	cnt := 0
	for i := 0; i < n; i++ {
		if ranges[i][0] > right {
			cnt++
		}
		right = max(right, ranges[i][1])
	}
	mod := int(math.Pow10(9)) + 7
	ans := 1
	for i := 0; i < cnt; i++ {
		ans = ans * 2 % mod
	}
	return ans
}
