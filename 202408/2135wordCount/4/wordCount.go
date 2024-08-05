package main

import (
	"fmt"
)

func main() {
	fmt.Println(wordCount([]string{"ant", "act", "tack"}, []string{"tack", "act", "acti"}))
}

// 在 startWords 或 targetWords 的任一字符串中，每个字母至多出现一次
func wordCount(startWords []string, targetWords []string) int {
	has := make(map[int]bool)

	for _, word := range startWords {
		c := 0
		for _, ch := range word {
			c = c | (1 << (int(ch) - 'a'))
		}
		has[c] = true
	}
	ans := 0
	for _, word := range targetWords {
		c := 0
		for _, ch := range word {
			c = c | (1 << (int(ch) - 'a'))
		}
		for _, ch := range word {
			b := c ^ (1 << (ch - 'a')) // 取点这个字符
			if has[b] {
				ans++
				break
			}
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
