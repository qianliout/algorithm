package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(minSwaps("100"))
}

func minSwaps(s string) int {
	a := strings.Count(s, "1")
	b := strings.Count(s, "0")
	if abs(a-b) > 1 {
		return -1
	}
	n := len(s)
	cnt1 := 0
	var pre byte = '0'
	for i := 0; i < n; i++ {
		if byte(s[i]) != pre {
			cnt1++
		}
		pre = revser(pre)
	}

	cnt2 := 0
	pre = '1'
	for i := 0; i < n; i++ {
		if byte(s[i]) != pre {
			cnt2++
		}
		pre = revser(pre)
	}
	if cnt1%2 != 0 {
		return cnt2 / 2
	} else if cnt2%2 != 0 {
		return cnt1 / 2
	}

	return min(cnt1, cnt2) / 2
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func revser(ch byte) byte {
	if ch == '0' {
		return '1'
	}
	return '0'
}
