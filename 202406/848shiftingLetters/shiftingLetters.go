package main

import (
	"fmt"
)

func main() {
	fmt.Println(shiftingLetters("abc", []int{3, 5, 9}))
}

func shiftingLetters(s string, shifts []int) string {
	n := len(s)
	suf := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suf[i] = suf[i+1] + shifts[i]
	}

	ss := []byte(s)

	ans := make([]byte, n)
	for i := range ss {
		ans[i] = byte((int(ss[i]-'a')+suf[i])%26 + 'a')
	}
	return string(ans)
}
