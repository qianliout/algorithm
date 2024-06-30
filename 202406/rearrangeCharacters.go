package main

import (
	"fmt"
)

func main() {
	fmt.Print()
}

func rearrangeCharacters(s string, target string) int {
	cnt1 := make(map[byte]int)
	for _, ch := range target {
		cnt1[byte(ch)]++
	}

	cnt2 := make(map[byte]int)
	for _, ch := range s {
		cnt2[byte(ch)]++
	}

	ans := len(s)
	for k, v := range cnt1 {
		ans = min(ans, cnt2[k]/v)
	}
	return ans
}

func check(cnt1, cnt2 map[byte]int) bool {
	for k, v := range cnt1 {
		if cnt2[k] < v {
			return false
		}
	}
	return true
}
