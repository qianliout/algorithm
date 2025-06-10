package main

import (
	"fmt"
)

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a."))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("aaa", "*"))
	fmt.Println(isMatch("aab", "*"))
	fmt.Println(isMatch("aaa", "a*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))

}

func isMatch(s string, p string) bool {
	return math([]byte(s), []byte(p), '0')
}

// 错的
func math(s, p []byte, pre byte) bool {
	if len(p) == 0 && len(s) == 0 {
		return true
	}
	if len(p) == 0 && len(s) != 0 {
		return false
	}
	if len(s) == 0 {
		if len(p) == 0 || (len(p) == 1 && p[0] == '*') {
			return true
		}
		return false
	}
	// 相同
	if s[0] == p[0] {
		return math(s[1:], p[1:], p[0])
	}

	if p[0] == '.' {
		return math(s[1:], p[1:], p[0])
	}
	if p[0] == '*' {
		ans := math(s, p[1:], '0')
		if pre == '.' {
			ans = ans || math(s[1:], p[1:], '0')
		}
		if pre != '0' && pre == s[0] {
			ans = ans || math(s[1:], p, pre)
		}
		return ans
	}
	return false
}

func match2(s, p []byte) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	// 这里的 len(s)>0 一定要记得，因为 s="",p="*"也是可以的
	first := len(s) > 0 && (s[0] == p[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		a := match2(s, p[2:])          // 匹配0个
		b := first && match2(s[1:], p) // 匹配多个
		return a || b
	}
	return first && match2(s[1:], p[1:])
}
