package main

import (
	"fmt"
)

func main() {
	fmt.Println(shortestBeautifulSubstring("100011001", 3))
	fmt.Println(shortestBeautifulSubstring("001110101101101111", 10))
	fmt.Println(Min("10101101101111", "11101011011011"))

}

func shortestBeautifulSubstring(s string, k int) string {
	ans := ""
	wind := 0
	ss := []byte(s)
	le, ri, n := 0, 0, len(s)

	for le <= ri && ri < n {
		if s[ri] == '1' {
			wind++
		}
		ri++
		for wind >= k {
			if wind == k {
				ans = Min(ans, string(ss[le:ri]))
			}
			if s[le] == '1' {
				wind--
			}
			le++
		}

	}
	return ans
}

func Min(pre, cur string) string {
	if pre == "" {
		return cur
	}
	if len(cur) > len(pre) {
		return pre
	}
	if len(pre) > len(cur) {
		return cur
	}
	if cur >= pre {
		return pre
	}
	return cur
}
