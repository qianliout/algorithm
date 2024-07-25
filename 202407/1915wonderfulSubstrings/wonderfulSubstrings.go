package main

import (
	"fmt"
)

func main() {
	fmt.Println(wonderfulSubstrings("aba"))
}

func wonderfulSubstrings(word string) int64 {
	cnt := make([]int, 1<<10)
	cnt[0] = 1 // 只有一个字母时
	ans, s := 0, 0
	for _, ch := range word {
		s = s ^ 1<<(int(ch)-int('a'))
		ans += cnt[s]
		for i := 0; i < 10; i++ {
			next := s ^ (1 << i)
			ans += cnt[next]
		}
		cnt[s]++
	}
	return int64(ans)
}
