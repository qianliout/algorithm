package main

import (
	"strings"
)

func main() {

}

func canBeTypedWords(text string, brokenLetters string) int {
	split := strings.Split(text, " ")
	cnt := make(map[byte]bool)
	for _, ch := range brokenLetters {
		cnt[byte(ch)] = true
	}
	ans := 0
	for _, ch := range split {
		if check(ch, cnt) {
			ans++
		}
	}

	return ans
}

func check(text string, cnt map[byte]bool) bool {
	for _, ch := range text {
		if cnt[byte(ch)] {
			return false
		}
	}
	return true
}
