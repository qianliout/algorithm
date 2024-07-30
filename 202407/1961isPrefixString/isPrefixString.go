package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPrefixString("ileetcode", []string{"i", "love", "leetcode", "apples"}))
}

func isPrefixString(s string, words []string) bool {
	for i := 1; i <= len(words); i++ {
		a := strings.Join(words[:i], "")
		// s和a 相等
		if s == a {
			return true
		}
	}
	return false
}
