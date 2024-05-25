package main

import (
	"fmt"
)

func main() {
	fmt.Println(smallestSubsequence("leetcode", 4, 'e', 2))
	fmt.Println(smallestSubsequence("leet", 3, 'e', 1))
	fmt.Println(smallestSubsequence("aaabbbcccddd", 3, 'b', 2))
	fmt.Println(smallestSubsequence("adffhjfmmmmorsfff", 6, 'f', 5))
	fmt.Println(smallestSubsequence("hjjhhhmhhwhz", 6, 'h', 5))
}

func smallestSubsequence(s string, k int, letter byte, repetition int) string {
	ss := []byte(s)
	win := make([]byte, 0)
	n := len(ss)
	exit := make(map[byte]int)
	add := make(map[byte]int)
	for _, ch := range ss {
		exit[ch]++
	}
	for i, ch := range ss {

		// pod := len(win) > 0
		// 没有那么

		for len(win) > 0 && ((
		// 和 letter 不同
		(win[len(win)-1] != letter && win[len(win)-1] >= ch && len(win)-1+n-i >= k) ||
			// 和letter 相同
			(win[len(win)-1] >= ch && win[len(win)-1] == letter && add[letter]-1+exit[letter] >= repetition && len(win)-1+n-i >= k)) ||
			// 没有那么多位置容纳接下来的数
			k-len(win) < (repetition-add[letter])) {

			add[win[len(win)-1]]--
			win = win[:len(win)-1]
		}
		win = append(win, ch)
		exit[ch]--
		add[ch]++
	}
	return string(win)
}
