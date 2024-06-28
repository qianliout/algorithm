package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestDupSubstring("banana"))
}

func longestDupSubstring(s string) string {
	n := len(s)
	ans := ""
	step := 0
	for i := 0; i < len(s); i++ {
		for j := i + step; j < n; j++ {
			sub := s[i : j+1]
			if strings.Contains(s[i+1:], sub) && len(sub) > len(ans) {
				ans = sub
			}
		}
		step = max(len(ans), 0)
	}
	return ans
}
