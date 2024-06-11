package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(merge("xyyyz", "xzyz"))
}

func minimumString(a string, b string, c string) string {
	ss := []string{a, b, c}

	perm := [][]int{{0, 1, 2}, {1, 0, 2}, {1, 2, 0}, {0, 2, 1}, {2, 0, 1}, {2, 1, 0}}
	ans := ""
	for _, ch := range perm {
		x, y, z := ch[0], ch[1], ch[2]
		s := merge(ss[x], merge(ss[y], ss[z]))
		if ans == "" || len(ans) > len(s) || (len(ans) == len(s) && s < ans) {
			ans = s
		}
	}
	return ans
}

func merge(s, t string) string {
	if strings.Contains(s, t) {
		return s
	}
	if strings.Contains(t, s) {
		return t
	}
	n := len(s)
	for i := min(len(s), len(t)); i >= 0; i-- {
		if s[n-i:] == t[:i] {
			return s + t[i:]
		}

	}
	return s + t
}
