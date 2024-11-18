package main

import (
	"fmt"
)

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	word := make(map[string]bool)
	for _, ch := range wordList {
		word[ch] = true
	}
	q := []string{beginWord}
	if !word[endWord] {
		return 0
	}
	ans := 0
	for len(q) > 0 {
		lve := make([]string, 0)
		for _, no := range q {
			if no == endWord {
				return ans + 1
			}
			nex := help(no, word)
			for _, ch := range nex {
				lve = append(lve, ch)
				word[ch] = false
			}
		}
		if len(lve) > 0 {
			ans++
		}
		q = lve
	}
	return 0
}

func help(s string, cnt map[string]bool) []string {
	ans := make([]string, 0)
	ss := []byte(s)
	for i := 0; i < len(ss); i++ {
		pre := ss[i]
		for j := 'a'; j <= 'z'; j++ {
			if pre == byte(j) {
				continue
			}
			ss[i] = byte(j)
			if cnt[string(ss)] {
				ans = append(ans, string(ss))
			}
		}
		ss[i] = pre
	}
	return ans
}
