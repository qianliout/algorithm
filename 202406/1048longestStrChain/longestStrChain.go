package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestStrChain([]string{"a", "b", "ba", "bca", "bda", "bdca"}))
}

func longestStrChain(words []string) int {
	m := make(map[string]bool)
	for _, ch := range words {
		m[ch] = true
	}
	ans := 0
	var dfs func(s string) int
	mem := make(map[string]int)
	dfs = func(s string) int {
		if !m[s] {
			return 0
		}
		ss := []byte(s)
		if va, ok := mem[s]; ok {
			return va
		}

		res := 1
		for i := 0; i < len(s); i++ {
			nx := append([]byte{}, ss[:i]...)
			nx = append(nx, ss[i+1:]...)
			if !m[string(nx)] {
				continue
			}
			res = max(res, dfs(string(nx))+1)
		}
		ans = max(ans, res)
		mem[s] = res
		return res
	}
	for i := range words {
		ans = max(ans, dfs(words[i]))
	}
	return ans
}
