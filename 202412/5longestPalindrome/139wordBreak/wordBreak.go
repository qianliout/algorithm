package main

import (
	"fmt"
)

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}

func wordBreak(s string, wordDict []string) bool {
	word := make(map[string]bool)
	for _, ch := range wordDict {
		word[ch] = true
	}
	n := len(s)
	f := make([]bool, n+1)
	f[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			f[i] = f[i] || (f[j] && word[s[j:i]])
		}
	}
	return f[n]
}
func wordBreak2(s string, wordDict []string) bool {
	word := make(map[string]bool)
	for _, ch := range wordDict {
		word[ch] = true
	}
	n := len(s)
	var dfs func(i int) bool
	mem := make([]int, n+10)
	dfs = func(i int) bool {
		if i <= 0 {
			return true
		}
		if mem[i] != 0 {
			return mem[i] == 1
		}
		ans := false
		for j := i - 1; j >= 0; j-- {
			ch := s[j:i]
			if word[ch] {
				ans = ans || dfs(j)
			}
		}
		if ans {
			mem[i] = 1
		} else {
			mem[i] = 2
		}
		return ans
	}
	ans := dfs(n)
	return ans
}
