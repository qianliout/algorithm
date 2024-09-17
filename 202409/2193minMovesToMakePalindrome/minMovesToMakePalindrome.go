package main

import (
	"fmt"
)

func main() {
	fmt.Println(minMovesToMakePalindrome("aabb"))
}

func minMovesToMakePalindrome(s string) int {
	ss := []byte(s)
	ans := 0
	for len(ss) > 0 {
		n := len(ss)
		c := ss[n-1]
		i := 0
		ss = ss[:n-1]
		find := false
		for ; i < n-1; i++ {
			if ss[i] == c {
				ans += i
				ss = append(ss[:i], ss[i+1:]...)
				find = true
				break
			}
		}
		if !find {
			ans += (n - 1) / 2
		}
	}
	return ans
}
