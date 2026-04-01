package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	ans, n := 0, len(s)
	win := make(map[byte]int)
	l, r := 0, 0
	for l <= r && r < n {
		c := s[r]
		win[c]++
		for win[c] > 1 {
			win[s[l]]--
			l++
		}
		r++
		ans = max(ans, r-l)
	}

	return ans
}
