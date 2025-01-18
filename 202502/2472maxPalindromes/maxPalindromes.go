package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxPalindromes2("abaccdbbd", 3))
	fmt.Println(maxPalindromes("abaccdbbd", 3))
	fmt.Println(maxPalindromes2("adbcda", 2))
	fmt.Println(maxPalindromes("adbcda", 2))
	fmt.Println(maxPalindromes2("kwnwkekokedadq", 5))
	fmt.Println(maxPalindromes("kwnwkekokedadq", 5))
}

func maxPalindromes(s string, k int) int {
	n := len(s)
	f := make([]int, n+1)
	ss := []byte(s)
	for i := 1; i <= n; i++ {
		f[i] = max(f[i], f[i-1]) // 这一步容易漏
		for j := i; j > 0; j-- {
			if len(ss[j-1:i]) >= k && check(ss[j-1:i]) {
				f[i] = max(f[i], f[j-1]+1)
			}
		}
	}
	return f[n]
}

func maxPalindromes2(s string, k int) int {
	n := len(s)
	ss := []byte(s)
	var dfs func(i int) int
	mem := make([]int, n+1)
	for i := range mem {
		mem[i] = n + 10
	}
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if mem[i] != n+10 {
			return mem[i]
		}
		ans := dfs(i - 1)

		// 选i 这个字符
		// 选一个长度至少是 k,且最后一个字符是 i 的回文串
		for j := i - k + 1; j >= 0; j-- {
			if check(ss[j : i+1]) {
				ans = max(ans, dfs(j-1)+1)
			}
		}
		mem[i] = ans
		return ans
	}
	ans := dfs(n - 1)
	return ans
}

func check(s []byte) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
