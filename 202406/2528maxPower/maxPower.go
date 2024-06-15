package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(maxPower([]int{4, 4, 4, 4}, 0, 3))
	fmt.Println(maxPower([]int{1, 2, 4, 5, 0}, 1, 2))
	fmt.Println(maxPower1([]int{1, 2, 4, 5, 0}, 1, 2))
}

func maxPower1(stations []int, r int, k int) int64 {
	n := len(stations)
	pre := make([]int, n+1)
	for i, ch := range stations {
		pre[i+1] = pre[i] + ch
	}
	start := make([]int, n) // 初始电量
	for i := 0; i < n; i++ {
		start[i] = pre[min(n, i+r+1)] - pre[max(0, i-r)]
	}

	var check func(mi int) bool

	// 一个问题，如果 城市i不够电量 mi,那么这个电站应该建在那里呢
	// 应该建在 min(n-1,i+r)处，也就是应该建立在右边，因为建在左边会有浪费
	// 这样这个电站的影响规范就是（i,mim(n-1,i+2r)）
	check = func(mi int) bool {
		d := make([]int, n)
		sumD := 0 // 累计差分
		need := 0 // 累计增加的电站
		for i := 0; i < n; i++ {
			sumD += d[i]
			m := mi - (start[i] + sumD)
			if m > 0 {
				need += m
				if need > k {
					return false
				}
				sumD += m

				if i+r*2+1 < n {
					d[i+r*2+1] -= m
				}
			}
		}
		return true
	}

	// 二分，右端点
	le := slices.Min(start)
	ri := le + k + 1
	for le < ri {
		mid := le + (ri-le+1)/2
		if check(mid) {
			le = mid
		} else {
			ri = mid - 1
		}
	}

	return int64(le)
}

func maxPower(stations []int, r int, k int) int64 {
	n := len(stations)
	sum := make([]int, n+1)
	for i, ch := range stations {
		sum[i+1] = sum[i] + ch
	}
	// 初始电量
	initP := make([]int, n)
	for i := range stations {
		initP[i] = sum[min(n, i+r+1)] - sum[max(0, i-r)]
	}

	// 意思是 最小的电量是 mi,那么新建的电站是否小于等于 k
	var check func(mi int) bool

	check = func(mi int) bool {
		d := make([]int, n)
		sumD := 0
		need := 0
		for i, ch := range initP {
			sumD += d[i]
			m := mi - (ch + sumD)
			if m > 0 {
				need += m
				if need > k {
					return false
				}
				sumD += m
				if i+2*r+1 < n {
					d[i+2*r+1] -= m
				}
			}
		}
		return true
	}
	// 开始二分
	a := slices.Min(initP)
	b := a + k + 1
	le, ri := a, b

	for le < ri {
		mid := le + (ri-le+1)/2
		if mid >= a && mid < b && check(mid) {
			le = mid
		} else {
			ri = mid - 1
		}
	}
	// 最后检测一下
	if le < a || le >= b || !check(le) {
		return -1
	}
	return int64(le)
}
