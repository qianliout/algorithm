package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestValidParentheses(")()())"))
}

func longestValidParentheses(s string) int {
	st := make([]int, 0)
	st = append(st, -1)
	ans := 0
	for i, c := range s {
		if c == '(' {
			st = append(st, i)
		} else {
			if len(st) > 0 {
				ans = max(ans, i-st[len(st)-1]+1)
				st = st[:len(st)]
			}
		}
	}
	return ans
}
