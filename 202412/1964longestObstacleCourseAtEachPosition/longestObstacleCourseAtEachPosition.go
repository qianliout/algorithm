package main

import (
	"math"
	"sort"
)

func main() {

}

func longestObstacleCourseAtEachPosition1(obstacles []int) []int {
	n := len(obstacles)
	inf := math.MaxInt64 / 100
	end := make([]int, n+1)
	for i := range end {
		end[i] = inf
	}
	end[0] = -inf
	ans := make([]int, n)

	mx := 1
	for i := 0; i < n; i++ {
		ch := obstacles[i]
		le, ri := 1, mx+1
		for le < ri {
			mid := le + (ri-le)/2
			if end[mid] > ch {
				ri = mid
			} else {
				le = mid + 1
			}
		}
		ans[i] = le
		end[le] = ch
		mx = max(mx, le)
	}
	return ans
}

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	ans := make([]int, len(obstacles))
	dp := make([]int, 0)
	for i, v := range obstacles {
		p := sort.SearchInts(dp, v+1) // 等价写法是二分不小于 v+1 的第一个位置
		if p < len(dp) {
			dp[p] = v
		} else {
			dp = append(dp, v)
		}
		ans[i] = p + 1
	}
	return ans
}
