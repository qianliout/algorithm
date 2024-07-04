package main

import (
	"strings"
)

func main() {

}

func reverseParentheses(s string) string {
	ss := []byte(s)

	st := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			st = append(st, i)
		} else if s[i] == ')' {
			// 不用判断，题目保证了是有效的
			le := st[len(st)-1] + 1
			st = st[:len(st)-1]
			ri := i - 1

			for le < ri {
				ss[le], ss[ri] = ss[ri], ss[le]
				le++
				ri--
			}
		}
	}
	res := string(ss)
	res = strings.ReplaceAll(res, ")", "")
	res = strings.ReplaceAll(res, "(", "")
	return res
}
