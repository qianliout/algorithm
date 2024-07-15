package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSubstrings("aba", "baba"))
}

func countSubstrings(s string, t string) int {
	m, n := len(s), len(t)
	mi := min(m, n)
	ans := 0
	for k := 1; k <= mi; k++ {
		for i := 0; i <= m-k; i++ {
			for j := 0; j <= n-k; j++ {
				cnt := cal(s[i:i+k], t[j:j+k])
				if cnt == 1 {
					ans++
				}
			}
		}
	}
	return ans
}

// 计算有多少个不同字符
func cal(s, t string) int {
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			cnt++
		}
	}
	return cnt
}
