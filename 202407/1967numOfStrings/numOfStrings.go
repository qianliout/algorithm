package main

import (
	"strings"
)

func main() {

}

func numOfStrings(patterns []string, word string) int {
	cnt := 0
	for _, ch := range patterns {
		if strings.Contains(word, ch) {
			cnt++
		}
	}
	return cnt
}
