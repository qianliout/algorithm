package main

import (
	"strings"
)

func main() {

}

func decodeString(s string) string {
	s2, _ := dfs([]byte(s), 0)
	return s2
}

func dfs(ss []byte, idx int) (string, int) {
	prefix, repeat := "", 0
	n := len(ss)
	for idx < n {
		c := ss[idx]
		if c >= '0' && c <= '9' {
			repeat = repeat*10 + int(c) - int('0')
		} else if c == '[' {
			tem, i := dfs(ss, idx+1)
			prefix += strings.Repeat(tem, repeat)
			idx = i
			repeat = 0
		} else if c == ']' {
			return prefix, idx
		} else {
			prefix = prefix + string(c)
		}
		idx++
	}
	return prefix, idx
}

/*
给定一个经过编码的字符串，返回它解码后的字符串。
编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
*/
