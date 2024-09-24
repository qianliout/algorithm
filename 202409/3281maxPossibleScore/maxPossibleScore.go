package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(maxPossibleScore([]int{2, 6, 13, 13}, 5))
}

func maxPossibleScore(start []int, d int) int {
	sort.Ints(start)
	n := len(start)
	// 最大的间距只能是在相邻两个区间中
	check := func(score int) bool {
		x := math.MinInt64 / 10
		for _, s := range start {
			x = max(x+score, s) // x 必须 >= 区间左端点 s
			if x > s+d {
				return false
			}
		}
		return true
	}

	mi, mx := 0, start[n-1]+d+1
	le, ri := mi, mx
	for le < ri {
		// 右端点
		mid := le + (ri-le+1)>>1
		if mid >= mi && mid < mx && check(mid) {
			le = mid
		} else {
			ri = mid - 1
		}
	}
	return le
}
