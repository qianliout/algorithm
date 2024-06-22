package main

import (
	"fmt"
)

func main() {
	fmt.Println(findAndReplacePattern([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb"))
}

func findAndReplacePattern(words []string, pattern string) []string {
	ans := make([]string, 0)
	for _, c := range words {
		if check(c, pattern) && check(pattern, c) {
			ans = append(ans, c)
		}
	}
	return ans
}

func check(word string, patten string) bool {
	cnt := make(map[byte]byte)
	if len(word) != len(patten) {
		return false
	}
	for i, c := range patten {
		nex := word[i]
		if cnt[byte(c)] == 0 {
			cnt[byte(c)] = byte(nex)
		} else if cnt[byte(c)] != byte(nex) {
			return false
		}
	}
	return true
}
