package main

import (
	"fmt"
)

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("cabwefgewcwaefgcf", "cae"))
}

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	tt := make(map[int]int, 26*2)
	for _, c := range t {
		tt[int(c)]++
	}
	ss := []byte(s)
	n := len(ss)
	window := make(map[int]int, 26*2)
	le, ri := 0, 0
	ans := ""
	for le <= ri && ri < n {
		c := int(ss[ri])
		window[c]++
		ri++
		for le <= ri && check(window, tt) {
			//  题目没有要求字典序最小
			// if ans == "" || len(ans) > len(ss[le:ri]) || ans > string(ss[le:ri]) {
			if ans == "" || len(ans) > len(ss[le:ri]) {
				ans = string(ss[le:ri])
			}
			ch := int(ss[le])
			window[ch]--
			le++
		}
	}
	return ans
}

func check(window, tt map[int]int) bool {
	for i, c := range tt {
		if window[i] < c {
			return false
		}
	}
	return true
}
