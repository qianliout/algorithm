package main

import (
	"fmt"
)

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}

func minWindow(s string, t string) string {
	win := make(map[byte]int)
	tt := make(map[byte]int)
	for i := range t {
		tt[t[i]]++
	}
	n := len(s)
	ans := ""
	left, right := 0, 0
	for left <= right && right < n {
		win[s[right]]++
		for check(win, tt) {
			if ans == "" || len(ans) > len(s[left:right+1]) {
				ans = s[left : right+1]
			}
			win[s[left]]--
			left++
		}
		right++
	}

	return ans
}

func check(a, b map[byte]int) bool {
	for k, v := range b {
		if a[k] < v {
			return false
		}
	}
	return true
}
