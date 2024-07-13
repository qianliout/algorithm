package main

import (
	"fmt"
)

func main() {
	fmt.Println(numSplits("aacaba"))
}

func numSplits(s string) int {
	cnt := make(map[byte]bool)
	n := len(s)
	pre := make([]int, n)
	pre[0] = 1
	cnt[byte(s[0])] = true
	for i := 1; i < n; i++ {
		ch := byte(s[i])
		pre[i] = pre[i-1]
		if !cnt[ch] {
			pre[i]++
			cnt[ch] = true
		}
	}
	cnt = make(map[byte]bool)
	cnt[byte(s[n-1])] = true
	suf := make([]int, n)
	suf[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		ch := byte(s[i])
		suf[i] = suf[i+1]
		if !cnt[ch] {
			suf[i]++
			cnt[ch] = true
		}
	}
	ans := 0
	for i := 0; i < n-1; i++ {
		if pre[i] == suf[i+1] {
			ans++
		}
	}
	return ans
}
