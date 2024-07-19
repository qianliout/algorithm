package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(minimumTeachings(2, [][]int{{1}, {2}, {1, 2}}, [][]int{{1, 2}, {1, 3}, {2, 3}}))
}

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	noCommon := make(map[int]bool)
	for _, ch := range friendships {
		if !hasCommon(languages, ch[0], ch[1]) {
			noCommon[ch[0]] = true
			noCommon[ch[1]] = true
		}
	}

	cnt := make([]int, n+1)

	for k := range noCommon {
		for _, ch := range languages[k-1] {
			cnt[ch]++
		}
	}
	mx := slices.Max(cnt)
	return len(noCommon) - mx
}

func hasCommon(lang [][]int, a, b int) bool {
	al := lang[a-1]
	bl := lang[b-1]
	for _, ch := range al {
		for _, ch2 := range bl {
			if ch == ch2 {
				return true
			}
		}
	}

	return false
}
