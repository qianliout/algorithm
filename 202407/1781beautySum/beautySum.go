package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	fmt.Println(beautySum("aabcb"))
}

func beautySum(s string) int {
	ans := 0
	n := len(s)
	for i := 0; i < n; i++ {
		cnt := make([]int, 26)
		for j := i; j < n; j++ {
			a := int(s[j]) - int('a')
			cnt[a]++
			// 不能直接用 slices.Min,因为有些字母的频率是0，也就是没有，需要把这部分滤过
			ans += slices.Max(cnt) - Min(cnt)
		}
	}
	return ans
}

func Min(cnt []int) int {
	ans := math.MaxInt / 10
	for _, v := range cnt {
		if v == 0 {
			continue
		}
		ans = min(ans, v)
	}
	return ans
}
