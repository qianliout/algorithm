package main

import (
	"fmt"
)

func main() {
	fmt.Println('a' - 'A')
	fmt.Println(longestNiceSubstring("YazaAay"))
	fmt.Println(longestNiceSubstring("dDzeE"))
	fmt.Println(longestNiceSubstring("SsZ"))
	// fmt.Println(longestNiceSubstring("cXlLwEZhtDPSiToVeWssVzzRMCrxmqSsZeIkzcWvzePMxsrpDMO"))
}

func longestNiceSubstring(s string) string {
	ans := ""
	n := len(s)
	for i := 0; i < n; i++ {
		cnt := make(map[byte]int)
		for j := i; j < n; j++ {
			c := byte(s[j])
			cnt[c]++
			if check(cnt) && (ans == "" || len(ans) < len(s[i:j+1])) {
				ans = s[i : j+1]
			}
		}
	}
	return ans
}

func check(cnt map[byte]int) bool {
	for k := range cnt {
		if k >= 'a' && k <= 'z' {
			if cnt[byte(k+'A'-'a')] <= 0 {
				return false
			}
		}
		if k >= 'A' && k <= 'Z' {
			if cnt[byte(k+'a'-'A')] <= 0 {
				return false
			}
		}
	}
	return true
}
