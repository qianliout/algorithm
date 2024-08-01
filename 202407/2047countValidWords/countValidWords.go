package main

import (
	"strings"
)

func main() {

}

func countValidWords(sentence string) int {
	split := strings.Split(sentence, " ")
	ans := 0
	for _, ch := range split {
		if check(ch) {
			ans++
		}
	}
	return ans
}

func check(s string) bool {
	s = strings.TrimSpace(s)
	b := 0
	if s == "" {
		return false
	}

	for i, ch := range s {
		if ch >= '0' && ch <= '9' {
			return false
		}
		if ch == '!' || ch == '.' || ch == ',' {
			return i == len(s)-1
		}
		if ch == '-' {
			if b > 0 {
				return false
			}
			b++
			if i > 0 && i < len(s)-1 && (s[i-1] >= 'a' && s[i-1] <= 'z') && (s[i+1] >= 'a' && s[i+1] <= 'z') {
				continue
			}
			return false

		}
	}
	return true
}
