package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(totalStrength([]int{5, 4, 6}))
	// fmt.Println(totalStrength([]int{1, 3, 1, 2}))
	fmt.Println(totalStrength([]int{558, 1567, 644, 40, 411, 1039, 37, 802, 461, 1001, 1814, 1195, 746,
		790, 374, 1215, 61, 1348, 14, 1501, 885, 885, 240, 1655, 1916, 763, 213, 1453, 281, 528, 1659, 938, 1266, 1109, 558, 808, 1086, 341, 920, 746}))
}

func totalStrength(strength []int) int {
	mod := int(math.Pow(10, 9)) + 7
	n := len(strength)
	left, right := make([]int, n), make([]int, n)
	preSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		left[i], right[i] = -1, n
		preSum[i+1] = preSum[i] + strength[i]
	}
	st := make([]int, 0)
	// 右边
	for i := 0; i < n; i++ {
		// right[i] 为右侧小于等于 strength[i] 的最近元素位置（不存在时为 n）
		for len(st) > 0 && strength[i] <= strength[st[len(st)-1]] {
			right[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	st = st[:0]
	// 左边
	for i := n - 1; i >= 0; i-- {
		// left[i] 为左侧严格小于 strength[i] 的最近元素位置（不存在时为 -1）
		for len(st) > 0 && strength[i] < strength[st[len(st)-1]] {
			left[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}

	// 这里没有太理解
	ans := 0
	s := 0                 // 前缀和
	ss := make([]int, n+2) // 前缀和的前缀和
	for i, v := range strength {
		s += v
		ss[i+2] = (ss[i+1] + s) % mod // 注意取模后，下面计算两个 ss 相减，结果可能为负
	}

	for i, v := range strength {
		l, r := left[i]+1, right[i]-1 // [l,r] 左闭右闭
		tot := ((i-l+1)*(ss[r+2]-ss[i+1]) - (r-i+1)*(ss[i+1]-ss[l])) % mod
		ans = (ans + v*tot) % mod // 累加贡献
	}

	return (ans + mod) % mod // 防止算出负数
}
