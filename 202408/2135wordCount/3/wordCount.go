package main

import (
	"fmt"
)

func main() {
	fmt.Println(wordCount([]string{"ant", "act", "tack"}, []string{"tack", "act", "acti"}))
}

// 在 startWords 或 targetWords 的任一字符串中，每个字母至多出现一次
func wordCount(startWords []string, targetWords []string) int {
	target := make([]int, len(targetWords))
	for i, word := range targetWords {
		c := 0
		for _, ch := range word {
			c = c | (1 << (int(ch) - 'a'))
		}
		target[i] = c
	}
	start := make(map[int]bool)
	for _, word := range startWords {
		ret := cal(word)
		for _, ch := range ret {
			start[ch] = true
		}
	}
	ans := 0
	for _, k := range target {
		if start[k] { // 这样写不可以，因为start 必须要操作一次
			ans++
		}
	}
	return ans
}

func cal(s string) []int {
	c := 0
	ans := make([]int, 0)
	for _, ch := range s {
		c = c | (1 << (int(ch) - 'a'))
	}
	for i := 0; i < 26; i++ {
		d := c | (1 << i)
		ans = append(ans, d)
	}
	return ans
}
