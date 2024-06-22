package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numSpecialEquivGroups([]string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"}))
}

func numSpecialEquivGroups(words []string) int {
	set := make(map[string]int)

	for _, word := range words {
		p := pire{cnt1: make([]int, 26), cnt2: make([]int, 26)}
		for i := 0; i < len(word); i += 2 {
			if i+1 < len(word) {
				p.cnt1[int(word[i+1])-int('a')]++
			}
			p.cnt2[int(word[i])-int('a')]++
		}
		key := p.key()
		set[key]++
	}
	return len(set)
}

type pire struct {
	cnt1 []int // 奇数
	cnt2 []int // 偶数
}

func (vi pire) key() string {
	ans := make([]string, 0)
	for i := 0; i < 26; i++ {
		ans = append(ans, fmt.Sprintf("%d", vi.cnt1[i]))
	}
	for i := 0; i < 26; i++ {
		ans = append(ans, fmt.Sprintf("%d", vi.cnt2[i]))
	}
	return strings.Join(ans, "")
}
