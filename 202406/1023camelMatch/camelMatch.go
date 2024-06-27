package main

import (
	"fmt"
)

func main() {
	fmt.Println(camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FB"))
}

func camelMatch(queries []string, pattern string) []bool {
	ans := make([]bool, len(queries))
	for i, ch := range queries {
		ans[i] = check(ch, pattern)
	}

	return ans
}

func check(s, p string) bool {
	m, n := len(s), len(p)
	l, r := 0, 0
	for r < n {
		for l < m && low(s[l]) && s[l] != p[r] {
			l++
		}
		if l == m || s[l] != p[r] {
			return false
		}
		l++
		r++
	}
	for l < m && low(s[l]) {
		l++
	}

	return l == m
}

func low(b byte) bool {
	if b >= 'a' && b <= 'z' {
		return true
	}
	return false
}
