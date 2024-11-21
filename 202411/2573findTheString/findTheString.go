package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(findTheString([][]int{{4, 0, 2, 0}, {0, 3, 0, 1}, {2, 0, 2, 0}, {0, 1, 0, 1}}))
}

func findTheString(lcp [][]int) string {
	n := len(lcp)
	ans := make([]byte, n)
	// 因为是字典序，所以从小到大，把能填的都填了
	for c := 'a'; c <= 'z'; c++ {
		// 找第一个还没有填的位置
		i := bytes.IndexByte(ans, 0)
		if i < 0 { // 说明所有的位置都有值了
			break
		}
		for j := i; j < n; j++ {
			if lcp[i][j] > 0 {
				ans[j] = byte(c)
			}
		}
	}

	for _, ch := range ans {
		if ch == 0 {
			return ""
		}
	}
	// 最后进行验证
	// 如果 s[i]=s[j]，那么 lcp[i][j]=lcp[i+1][j+1]+1；
	// 如果 s[i]!=s[j]，那么 lcp[i][j]=0。
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			cl := 0
			if ans[i] == ans[j] {
				cl = 1
				if i < n-1 && j < n-1 {
					cl = lcp[i+1][j+1] + 1
				}
			}
			if lcp[i][j] != cl {
				return ""
			}
		}
	}

	return string(ans)
}

// 错误的解法
func findTheString2(lcp [][]int) string {
	n := len(lcp)
	ans := make([]byte, n)

	for i := 0; i < n; i++ {
		if ans[i] != 0 {
			continue
		}
		c := byte('a')
		if i > 0 {
			if ans[i-1] == 'z' {
				return ""
			}
			c = byte(ans[i-1] + 1)
		}
		ans[i] = c
		for j := i + 1; j < n; j++ {
			if lcp[i][j] > 0 {
				ans[j] = c
			}
		}
	}
	for _, ch := range ans {
		if ch == 0 {
			return ""
		}
	}
	// 最后进行验证
	// 如果 s[i]=s[j]，那么 lcp[i][j]=lcp[i+1][j+1]+1；
	// 如果 s[i]!=s[j]，那么 lcp[i][j]=0。

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			cl := 0
			if i+1 < n && j+1 < n && ans[i] == ans[j] {
				cl = lcp[i+1][j+1] + 1
			}
			if i == n-1 || n == n-1 && ans[i] == ans[j] {
				cl = 1
			}
			if lcp[i][j] != cl {
				return ""
			}
		}
	}

	return string(ans)
}
