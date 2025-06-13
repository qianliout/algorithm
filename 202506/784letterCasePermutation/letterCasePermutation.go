package main

import (
	"unicode"
)

func main() {

}

func letterCasePermutation(s string) []string {
	ans := make([]string, 0)
	used := make(map[string]bool)
	ss := []byte(s)
	var dfs func(start int)
	n := len(s)
	dfs = func(start int) {
		if !used[string(ss)] {
			used[string(ss)] = true
			ans = append(ans, string(ss))
		}
		if start < 0 || start >= n {
			return
		}
		for i := start; i < n; i++ {
			if ss[i] >= '0' && ss[i] <= '9' {
				continue
			}
			pre := ss[i]
			ss[i] = byte(swapCase(rune(pre)))
			dfs(i + 1)
			ss[i] = pre
		}
	}
	dfs(0)
	return ans
}

func swapCase(c rune) rune {
	if unicode.IsLower(c) {
		return unicode.ToUpper(c)
	} else if unicode.IsUpper(c) {
		return unicode.ToLower(c)
	}
	return c // 非字母字符直接返回
}
