package main

import (
	"fmt"
)

func main() {
	// fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "*"))
	// fmt.Println(isMatch("cb", "?a"))
}

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	f := make([][]bool, m+1)
	for i := range f {
		f[i] = make([]bool, n+1)
	}
	// 初值
	for i := 0; i <= m; i++ {
		f[i][0] = false
	}
	for j := 0; j <= n; j++ {
		f[0][j] = false
	}
	f[0][0] = true

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i-1][j-1] || f[i][j-1] || f[i-1][j]
			}
		}
	}
	return f[m][n]
}

// '?' 可以匹配任何单个字符。
// '*' 可以匹配任意字符序列（包括空字符序列）。
// 第10的条件
// 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
// '.' 匹配任意单个字符
// '*' 匹配零个或多个前面的那一个元素
