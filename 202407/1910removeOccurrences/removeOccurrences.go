package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(removeOccurrences("aabababa", "aba"))
}

func removeOccurrences(s string, part string) string {
	for {
		// 这样写会出错:aabababa
		// next := strings.ReplaceAll(s, part, "")
		next := strings.Replace(s, part, "", 1)
		if next == s {
			return s
		}
		s = next
	}
}
