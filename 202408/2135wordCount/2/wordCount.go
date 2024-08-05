package main

import (
	"fmt"
)

func main() {
	fmt.Println(wordCount([]string{"ant", "act", "tack"}, []string{"tack", "act", "acti"}))
}

// 在 startWords 或 targetWords 的任一字符串中，每个字母至多出现一次

// 解法不正确
func wordCount(startWords []string, targetWords []string) int {
	target := make(map[string]int)
	for _, ch := range targetWords {
		target[ch]++
	}
	ans := 0
	used := make(map[string]int)
	for _, ch := range startWords {
		ret := cal(ch)
		for _, w := range ret {
			if target[w] > 0 && used[w] == 0 {
				ans++
				used[w]++
			}
		}
	}
	return ans
}

func cal(s string) []string {
	// 增加一个
	ans := make([]string, 0)
	for i := 'a'; i <= 'z'; i++ {
		a := s + string(i)
		ans = append(ans, a)
	}
	n := len(s)
	used := make([]bool, n)
	ss := []byte(s)
	var dfs func(path []byte)
	dfs = func(path []byte) {
		if len(path) == len(s) {
			ans = append(ans, string(path))
		}
		for i := 0; i < n; i++ {
			if !used[i] {
				used[i] = true
				path = append(path, ss[i])
				dfs(path)
				used[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs([]byte{})
	return ans
}
