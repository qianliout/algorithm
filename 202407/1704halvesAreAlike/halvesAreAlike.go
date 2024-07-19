package main

import (
	"strings"
)

func main() {

}

func halvesAreAlike(s string) bool {
	s = strings.ToLower(s)
	n := len(s)
	pre := s[:n/2]
	nex := s[n/2:]
	cnt1, cnt2 := 0, 0
	for _, ch := range pre {
		switch byte(ch) {
		case 'a', 'e', 'i', 'o', 'u':
			cnt1++
		}
	}

	for _, ch := range nex {
		switch byte(ch) {
		case 'a', 'e', 'i', 'o', 'u':
			cnt2++
		}
	}

	return cnt1 == cnt2
}
