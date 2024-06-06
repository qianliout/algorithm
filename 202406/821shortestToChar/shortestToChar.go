package main

import (
	"fmt"
)

func main() {
	fmt.Println(shortestToChar("loveleetcode", 'e'))
	fmt.Println(shortestToChar("aaab", 'b'))
}

func shortestToChar(s string, c byte) []int {
	ss := []byte(s)
	n := len(s)
	ans := make([]int, len(s))
	for i := 0; i < n; i++ {
		if ss[i] == c {
			ans[i] = 0
			continue
		}
		// left
		ans[i] = n
		for j := i - 1; j >= 0; j-- {
			if ss[j] == c {
				ans[i] = min(ans[i], i-j)
			}
		}
		for j := i + 1; j < n; j++ {
			if ss[j] == c {
				ans[i] = min(ans[i], j-i)
			}
		}
	}

	return ans
}
