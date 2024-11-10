package main

import (
	"fmt"
)

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

func minWindow(s string, t string) string {
	tt := make(map[byte]int)
	for _, ch := range t {
		tt[byte(ch)]++
	}
	wind := make(map[byte]int)
	left, right := 0, 0
	ans := ""
	for left <= right && right < len(s) {
		wind[s[right]]++
		right++
		for left <= right && check(wind, tt) {
			// 保证了唯一
			if ans == "" || len(ans) > right-left || (len(ans) == right-left && ans > s[left:right]) {
				ans = s[left:right]
			}
			wind[s[left]]--
			left++
		}
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
