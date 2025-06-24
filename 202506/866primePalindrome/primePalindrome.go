package main

import (
	"fmt"
	"math"
)

func main() {

}

func primePalindrome(n int) int {
	for i := n; i < math.MaxInt64/2; i++ {
		if check(i) {
			return i
		}
	}
	return 0
}

// 会超时
func check(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	res := fmt.Sprintf("%d", n)
	l, r := 0, len(res)-1
	for l < r {
		if res[l] != res[r] {
			return false
		}
		l++
		r--
	}
	return true
}
