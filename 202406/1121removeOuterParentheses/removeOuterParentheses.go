package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))
}

func removeOuterParentheses2(s string) string {
	cnt := 0
	ans := make([]byte, 0)
	for _, ch := range s {
		if byte(ch) == '(' {
			if cnt > 0 {
				ans = append(ans, '(')
			}
			cnt++
		}
		if byte(ch) == ')' {
			if cnt > 1 {
				ans = append(ans, ')')
			}
			cnt--
		}
	}
	return string(ans)
}

func removeOuterParentheses(s string) string {
	// 怎么判断一个（ 是应该保留的，那就是这个（ 外面还有（ ， 怎么判断外面有还有（ 呢，可以用记数也可以用栈的方式
	// 怎么判断一个) 是应该保留的呢，因为这佧字符串是有效的括号字符串，所有 外面有 ），那前面一定有多余的 (,
	// 题目只要求删除最外层的括号，那个字符串只有一层呢，题目中也要删除
	// 也就是说只保留有外层括号的括号
	stark := make([]byte, 0)
	ans := make([]byte, 0)
	for _, ch := range s {
		if ch == '(' {
			if len(stark) >= 1 {
				ans = append(ans, '(')
			}
			stark = append(stark, '(')
		}
		if ch == ')' {
			if len(stark) >= 2 {
				ans = append(ans, ')')
			}
			stark = stark[:len(stark)-1]
		}
	}
	return string(ans)
}
