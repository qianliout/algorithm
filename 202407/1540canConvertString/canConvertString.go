package main

import (
	"fmt"
)

func main() {
	// fmt.Println(canConvertString("input", "ouput", 9))
	// fmt.Println(canConvertString("abc", "bcd", 10))
	fmt.Println(canConvertString("iqssxdlb", "dyuqrwyr", 40))
}

// 文字游戏啊
func canConvertString(s string, t string, k int) bool {
	if len(s) != len(t) {
		return false
	}
	ans := make([]int, 26)
	n := len(s)

	for i := 0; i < n; i++ {
		x := int(t[i]) - int(s[i])
		// 这样写是不可以的
		// x := int(t[i] - s[i])
		if x < 0 {
			x += 26
		}
		// 这个值在上一轮用过，那么就只能等到下一轮了
		if ans[x] > 0 {
			ans[x] += 26
		} else {
			ans[x] = x
		}
		if ans[x] > k {
			return false
		}
	}
	return true
}
