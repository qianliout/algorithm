package main

import (
	"fmt"
)

func main() {
	fmt.Println(canConstruct("annabelle", 2))
}

// 如果要求是子串，那么这个题这样做是对的，
func canConstruct(s string, k int) bool {
	n := len(s)
	var dfs func(j, k int) bool
	dfs = func(j, k int) bool {
		if k == 1 {
			return check(s, 0, j)
		}
		res := false
		for m := j; m >= 0; m-- {
			if check(s, m, j) {
				res = res || dfs(m-1, k-1)
			}
		}

		return res
	}
	res := dfs(n-1, k)
	return res
}

func check(s string, i, j int) bool {
	if i > j {
		return false
	}
	le, ri := i, j
	for le < ri {
		if s[le] != s[ri] {
			return false
		}
		le++
		ri--
	}
	return true
}
