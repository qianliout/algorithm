package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumCostSubstring("adaa", "d", []int{-1000}))
}

// 数据范围：-1000 <= vals[i] <= 1000
func maximumCostSubstring(s string, chars string, vals []int) int {
	vals2 := make([]int, 26)
	for i := 0; i < 26; i++ {
		vals2[i] = i + 1 // 值从1开始
	}
	for i, ch := range chars {
		idx := int(ch) - int('a')
		vals2[idx] = vals[i]
	}

	n := len(s)
	f := make([]int, n+1)
	ans := 0
	for i, ch := range s {
		idx := int(ch) - int('a')
		v := vals2[idx]
		f[i+1] = max(f[i]+v, v)
		ans = max(ans, f[i+1])
	}
	return ans
}
