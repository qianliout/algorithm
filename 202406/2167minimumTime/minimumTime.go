package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumTime("1100101"))
	fmt.Println(minimumTime("0010"))
	fmt.Println(minimumTime("011001111111101001010000001010011"))
	fmt.Println(minimumTime("110001110000100001100010111101010011101101000111"))
}

func minimumTime1(s string) int {
	n := len(s)
	// pre[i] 表示s[:i]的结果，不包括 i
	pre := make([]int, n+1)
	// suf[i]表示s[i:]后面的结果值
	suf := make([]int, n+1)
	/*
		当 s[i]=0 时，无需移除车厢，则有 pre[i]=pre[i−1]；
		当 s[i]=1时，可以单独移除第 i 节车厢，也可以移除前 i 个车厢，二者取最小值。
	*/
	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			suf[i] = suf[i+1]
		} else {
			suf[i] = min(suf[i+1]+2, n-i)
		}
	}
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			pre[i+1] = pre[i]
		} else {
			pre[i+1] = min(pre[i]+2, i+1)
		}
	}

	ans := n
	for i := 0; i < n; i++ {
		ans = min(ans, pre[i+1]+suf[i+1])
	}
	return ans
}
func minimumTime(s string) int {
	n := len(s)
	suf := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			suf[i] = suf[i+1]
		} else {
			suf[i] = min(suf[i+1]+2, n-i)
		}
	}
	ans := suf[0]
	pre := 0
	for i, ch := range s {
		if ch == '1' {
			pre = min(pre+2, i+1)
			ans = min(ans, pre+suf[i+1])
		}
	}
	return ans
}
