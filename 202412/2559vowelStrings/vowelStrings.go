package main

import (
	"strings"
)

func main() {

}
func vowelStrings(words []string, queries [][]int) []int {
	n := len(words)
	pre := make([]int, n+1)
	for i, ch := range words {
		pre[i+1] = check(ch) + pre[i]
	}
	ans := make([]int, len(queries))
	for i, ch := range queries {
		l, r := ch[0], ch[1]
		ans[i] = pre[r+1] - pre[l]
	}

	return ans

}

func check(word string) int {
	start := word[0]
	end := word[len(word)-1]
	if strings.Contains("aeiou", string(start)) && strings.Contains("aeiou", string(end)) {
		return 1
	}
	return 0
}
