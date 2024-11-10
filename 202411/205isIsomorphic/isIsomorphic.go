package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isIsomorphic("badc", "baba"))
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt1 := make(map[byte]int)
	cnt2 := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		cnt1[s[i]] = i
		cnt2[t[i]] = i
	}
	for k := range s {
		if cnt1[s[k]] != cnt2[t[k]] {
			return false
		}
	}
	return true
}

func wordPattern(pattern string, s string) bool {
	cnt1 := make(map[byte]int)
	cnt2 := make(map[string]int)
	ss := strings.Split(s, " ")
	if len(pattern) != len(ss) {
		return false
	}
	for i := range pattern {
		cnt1[byte(pattern[i])] = i
		cnt2[ss[i]] = i
	}
	for i := range pattern {
		if cnt1[byte(pattern[i])] != cnt2[ss[i]] {
			return false
		}
	}
	return true
}
