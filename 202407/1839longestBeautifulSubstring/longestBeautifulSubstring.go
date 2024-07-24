package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestBeautifulSubstring("aeiaaioaaaaeiiiiouuuooaauuaeiu"))
	fmt.Println(longestBeautifulSubstring("a"))
}

var obs []byte = []byte{'a', 'e', 'i', 'o', 'u'}

// word 只包含字符 'a'，'e'，'i'，'o' 和 'u' 。
func longestBeautifulSubstring(word string) int {
	left := 0
	right := 0
	n := len(word)
	ans := 0
	for ; right < n; right++ {
		// 从a 开头
		if byte(word[left]) != 'a' {
			left++
			continue
		}
		i := 0
		for ; i < 5; i++ {
			if right >= n {
				break
			}
			cur := byte(obs[i])
			for right < n && cur == byte(word[right]) {
				right++
			}
			if i < 4 && right < n && word[right] != obs[i+1] {
				break
			}
		}
		if i == 5 {
			ans = max(ans, right-left)
		}
		left = right
	}
	return ans
}
