package main

func main() {

}

func isMatch(s string, p string) bool {
	s = " " + s
	p = " " + p
	return match([]byte(s), []byte(p))
}

/*
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s 的，而不是部分字符串。
*/

func match(s, p []byte) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	first := len(s) > 0 && (s[0] == p[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		a := match(s, p[2:])
		b := first && match(s[1:], p)
		return a || b
	}
	return first && match(s[1:], p[1:])
}
