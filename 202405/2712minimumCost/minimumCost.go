package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumCost("010101"))
}

func minimumCost(s string) int64 {
	ans := 0
	n := len(s)
	for i := 0; i < n-1; i++ {
		if s[i] != s[i+1] {
			ans += min(i+1, n-i-1)
		}
	}
	return int64(ans)
}
