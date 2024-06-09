package main

import (
	"fmt"
)

func main() {
	fmt.Println(minFlips("111000"))
	fmt.Println(minFlips("0001100010101000111101000110101111000000101100000001001"))
}

// 很朴素的做法，但是会超时
func minFlips(s string) int {
	n := len(s)
	s = s + s
	ss := []byte(s)
	le, ri := 0, n
	ans := n
	for ri <= 2*n {
		ans = min(ans, cac(ss[le:ri]))
		le++
		ri++
	}
	return ans
}

func cac(by []byte) int {
	ans1, ans2 := 0, 0
	for i := 0; i < len(by); i++ {
		if int(by[i]-'0') != i%2 {
			ans1++
		}
		if int(by[i]-'0') != (i+1)%2 {
			ans2++
		}
	}
	return min(ans1, ans2)
}
