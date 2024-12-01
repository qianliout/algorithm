package main

import (
	"strings"
)

func main() {

}

func vowelStrings(words []string, left int, right int) int {
	ans := 0
	n := len(words)
	for i := max(0, left); i <= min(right, n-1); i++ {
		ch := words[i]
		if strings.Contains("aeiou", string(ch[0])) &&
			strings.Contains("aeiou", string(ch[len(ch)-1])) {
			ans++
		}
	}
	return ans
}
