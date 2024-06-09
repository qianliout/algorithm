package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumScore("abacaba", "bzaa"))
}

/*
定义 pre[i]\textit{pre}[i]pre[i] 为 s[:i]s[:i]s[:i] 对应的 ttt 的最长前缀的结束下标。

定义 suf[i]\textit{suf}[i]suf[i] 为 s[i:]s[i:]s[i:] 对应的 ttt 的最长后缀的开始下标
*/
func minimumScore(s string, t string) int {
	n, m := len(s), len(t)
	suf := make([]int, n+1)
	for i := range suf {
		suf[i] = m
	}
	j := m - 1

	for i := n - 1; i >= 0; i-- {
		if j >= 0 && s[i] == t[j] {
			j--
		}
		suf[i] = j + 1
	}
	ans := suf[0]
	if ans == 0 {
		return ans
	}
	j = 0
	for i := range s {
		if s[i] == t[j] {
			ans = min(ans, suf[i+1]-j-1)
			j++
		}
	}
	return ans
}
