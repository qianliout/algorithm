package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumMoves("XXX"))
}

func minimumMoves(s string) int {
	ans := 0
	n := len(s)
	for i := 0; i < n; {
		if s[i] == 'O' {
			i++
			continue
		}
		ans++
		i = i + 3
	}
	return ans
}
