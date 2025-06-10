package main

import (
	"fmt"
)

func main() {
	fmt.Println(isMatch("aab", "a*?b"))
	fmt.Println(isMatch("aab", "a?*b"))
	fmt.Println(isMatch("aab", "aa*b"))
	fmt.Println(isMatch("aab", "aa?b"))
	fmt.Println(isMatch("", "*******"))
	fmt.Println(isMatch("abbabbbaabaaabbbbbabbabbabbbabbaaabbbababbabaaabbab", "*aabb***aa**a******aa*"))

}
func isMatch(s string, p string) bool {
	mem := make(map[string]map[string]bool)
	// return match2([]byte(s), []byte(p))
	return match3(s, p, mem)
}

// 可以得到结果，但是会超时
func match2(s, p []byte) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	// 这里的 len(s)>0 一定要记得，因为 s="",p="*"也是可以的
	first := len(s) > 0 && (s[0] == p[0] || p[0] == '?')
	if p[0] == '*' {
		a := match2(s, p[1:])
		if len(s) > 0 {
			b := match2(s[1:], p[1:])
			c := match2(s[1:], p)
			a = a || b || c
		}
		return a
	}
	return first && match2(s[1:], p[1:])
}

// 加上缓存能过
func match3(s, p string, mem map[string]map[string]bool) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	if s != "" {
		if mem[s] != nil {
			v, ok := mem[s][p]
			if ok {
				return v
			}
		}
	}

	// 这里的 len(s)>0 一定要记得，因为 s="",p="*"也是可以的
	first := len(s) > 0 && (s[0] == p[0] || p[0] == '?')
	if p[0] == '*' {
		a := match3(s, p[1:], mem)
		if len(s) > 0 {
			b := match3(s[1:], p[1:], mem)
			c := match3(s[1:], p, mem)
			a = a || b || c
		}
		if mem[s] == nil {
			mem[s] = make(map[string]bool)
		}
		mem[s][p] = a
		return a
	}
	ans := first && match3(s[1:], p[1:], mem)
	if mem[s] == nil {
		mem[s] = make(map[string]bool)
	}
	mem[s][p] = ans
	return ans
}

/*
'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符序列（包括空字符序列）。
*/
/*
'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
*/

func isMatch2(s string, p string) bool {
	m, n := len(s), len(p)
	//  f[i][j] 表示前 i,j 个字符的结果，不包括 i,j
	f := make([][]bool, m+1)
	for i := range f {
		f[i] = make([]bool, n+1)
	}
	// 初值
	// 这是这个题目的关键
	f[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] != '*' {
			break
		}
		f[0][i] = true
	}

	for j := 1; j <= m; j++ {
		f[j][0] = false // 因为初值是 false，可以不写
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-1] || f[i-1][j]
			}
		}
	}
	return f[m][n]
}

func isMatch3(s string, p string) bool {
	// s:="" p = "*****"
	s = " " + s
	p = " " + p

	m, n := len(s), len(p)
	f := make([][]bool, m+5)
	for i := range f {
		f[i] = make([]bool, n+5)
	}
	f[0][0] = true
	//  其他情况下的初值都是false
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				f[i][j] = f[i-1][j-1]
			}
			if p[j-1] == '*' {
				// 可以0个或多个，那就向前找
				for k := i; k > 0; k-- {
					if f[k][j-1] {
						f[i][j] = true
						break
					}
				}
			}
		}
	}
	return f[m][n]
}
