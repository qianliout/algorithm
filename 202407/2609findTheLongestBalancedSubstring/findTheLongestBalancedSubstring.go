package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findTheLongestBalancedSubstring("01000111"))
	fmt.Println(findTheLongestBalancedSubstring("111"))
}

func findTheLongestBalancedSubstring(s string) int {
	ans, n := 0, len(s)
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if check(s[i:j]) {
				ans = max(ans, j-i)
			}
		}
	}
	return ans
}

func check(s string) bool {
	n := len(s)
	fir := strings.Index(s, "1")
	if fir == -1 {
		return false
	}
	i := fir
	for i < n {
		if s[i] != '1' {
			break
		}
		i++
	}
	return i == n && n-fir == fir
}
