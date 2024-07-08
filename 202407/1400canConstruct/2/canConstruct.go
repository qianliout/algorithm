package main

import (
	"fmt"
)

func main() {
	fmt.Println(canConstruct("annabelle", 2))
}

// 题目没有要求子串
func canConstruct(s string, k int) bool {
	table := make([]int, 26)
	for _, ch := range s {
		table[int(ch-'a')]++
	}
	cnt := 0
	for _, ch := range table {
		if ch&1 == 1 {
			cnt++
		}
	}
	return cnt <= k && k <= len(s)
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
