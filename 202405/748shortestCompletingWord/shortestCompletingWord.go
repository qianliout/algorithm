package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(shortestCompletingWord("GrC8950", []string{"measure", "other", "every", "base", "according", "level", "meeting", "none", "marriage", "rest"}))
}

func shortestCompletingWord(licensePlate string, words []string) string {

	licensePlate = strings.ToLower(licensePlate)
	licens := make([]int, 26)
	for _, ch := range licensePlate {
		if ch >= 'a' && ch <= 'z' {
			licens[ch-'a']++
		}
	}
	ans := ""
	for _, word := range words {
		if same(licens, word) {
			if ans == "" || len(ans) > len(word) {
				ans = word
			}
		}
	}
	return ans
}

func same(a []int, b string) bool {
	b = strings.ToLower(b)
	he := make([]int, 26)
	for _, ch := range b {
		if ch >= 'a' && ch <= 'z' {
			he[ch-'a']++
		}
	}
	for i := 0; i < 26; i++ {
		if a[i] == 0 {
			continue
		}

		if a[i] > he[i] {
			return false
		}
	}
	return true
}
