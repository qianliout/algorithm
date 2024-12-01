package main

import (
	"fmt"
)

func main() {
	fmt.Println(ladderLength("a", "c", []string{"a", "b", "c"}))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 0
	}
	word := make(map[string]bool)
	for _, ch := range wordList {
		word[ch] = true
	}
	if !word[endWord] {
		return 0
	}
	q := []string{beginWord}
	ans := 0
	for len(q) > 0 {
		lev := make([]string, 0)
		for _, no := range q {
			if no == endWord {
				return ans + 1
			}

			next := help(no, word)
			for _, nx := range next {
				if word[nx] {
					lev = append(lev, nx)
					word[nx] = false
				}
			}
		}
		if len(lev) > 0 {
			ans++
		}
		q = lev
	}
	return 0
}

func help(start string, bank2 map[string]bool) []string {
	ans := make([]string, 0)
	ss := []byte(start)
	n := len(ss)
	for i := 0; i < n; i++ {
		pre := ss[i]
		for j := 'a'; j <= 'z'; j++ {
			if pre == byte(j) {
				continue
			}
			ss[i] = byte(j)
			// if !bank2[string(ss)] {
			// 	continue
			// }
			ans = append(ans, string(ss))
		}
		ss[i] = pre
	}
	return ans
}
