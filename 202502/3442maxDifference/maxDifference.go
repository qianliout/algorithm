package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxDifference("aaaaabbc"))
	fmt.Println(maxDifference("tzt"))
}

func maxDifference(s string) int {
	cnt := make([]int, 26)
	for _, ch := range s {
		cnt[ch-'a']++
	}
	n := len(s)
	a := 0 // 出现奇数次的最大
	b := n // 出现偶数次的最小
	for i := 0; i < 26; i++ {
		if cnt[i] == 0 {
			continue
		}
		if cnt[i]%2 == 0 {
			b = min(b, cnt[i])
		} else {
			a = max(a, cnt[i])
		}
	}
	return a - b
}
