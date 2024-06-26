package main

import (
	"fmt"
)

func main() {
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))
}

func commonChars(words []string) []string {
	n := len(words)
	cnt := make([][26]int, n)
	for i, word := range words {
		for _, ch := range word {
			idx := int(ch) - int('a')
			cnt[i][idx]++
		}
	}
	res := make([]string, 0)

	for j := 0; j < 26; j++ {
		all := n + 1
		for i := 0; i < n; i++ {
			all = min(all, cnt[i][j])
		}
		for i := 0; i < all; i++ {
			res = append(res, string(byte(j+'a')))
		}
	}

	return res
}
