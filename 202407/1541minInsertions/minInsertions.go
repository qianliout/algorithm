package main

import (
	"fmt"
)

func main() {
	fmt.Println(minInsertions("())"))
	fmt.Println(minInsertions("))())("))
	fmt.Println(minInsertions("(()))"))
	fmt.Println(minInsertions("(()))(()))()())))"))
}

/*
给你一个括号字符串 s ，它只包含字符 '(' 和 ')' 。一个括号字符串被称为平衡的当它满足：

	任何左括号 '(' 必须对应两个连续的右括号 '))' 。 注意这里不一定是挨着
	左括号 '(' 必须在对应的连续两个右括号 '))' 之前。
*/
func minInsertions2(s string) int {
	ans := 0
	left := 0
	n := len(s)
	i := 0
	for i < n {
		if s[i] == '(' {
			left++
		} else {
			if i+1 < n && s[i+1] == ')' {
				i++
			} else {
				ans++
			}
			left--
			if left < 0 {
				ans++
				left++
			}
		}
		i++
	}
	return ans + left*2 //   多出的左括号都匹配两个右括号
}

func minInsertions(s string) int {
	ans := 0
	left := 0
	n := len(s)
	i := 0
	for ; i < n; i++ {
		if s[i] == '(' {
			left++
			continue
		}
		// s[i]和s[i+1]都是 )
		if i+1 < n && s[i+1] == ')' {
			i++
		} else {
			ans++
		}

		left--
		if left < 0 {
			ans++
			left++
		}
	}
	return ans + left*2 //   多出的左括号都匹配两个右括号
}
