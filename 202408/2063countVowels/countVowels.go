package main

import (
	"strings"
)

func main() {

}

func countVowels(word string) int64 {
	n := len(word)
	var ans int64
	for i, ch := range word {
		if strings.ContainsRune("aeiou", ch) {
			l := i + 1
			r := n - i
			ans += int64(l) * int64(r)
		}
	}
	return ans
}
