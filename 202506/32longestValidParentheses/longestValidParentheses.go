package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestValidParentheses(")()())"))
}

func longestValidParentheses(s string) int {
	n := len(s)
	f := make([]int, n)
	st := make([]int, 0)
	ans := 0
	for i, c := range s {
		if c == '(' {
			st = append(st, i)
		} else {
			if len(st) > 0 {
				pre := st[len(st)-1]
				th := i - pre + 1
				f[i] = th
				if pre-1 >= 0 {
					f[i] += f[pre-1]
				}
				st = st[:len(st)-1]
				ans = max(ans, f[i])
			}
		}
	}
	return ans
}
