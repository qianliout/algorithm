package main

import (
	"fmt"
)

func main() {
	fmt.Println(minDeletions("aaabbbcc"))
}

func minDeletions(s string) int {
	set := make(map[int]bool)
	cnt := make([]int, 26)
	for _, ch := range s {
		idx := int(ch) - int('a')
		cnt[idx]++
	}
	ans := 0
	for _, v := range cnt {
		if v == 0 {
			continue
		}
		for set[v] {
			v--
			ans++
		}
		if v != 0 {
			set[v] = true
		}

	}

	return ans
}
