package main

import (
	"strings"
)

func main() {

}

func maximumSubsequenceCount(text string, pattern string) int64 {
	a, b := pattern[0], pattern[1]
	if a == b {
		c := strings.Count(text, string(a))
		return int64(c * (c + 1) / 2)
	}

	ac, bc := 0, 0
	ans := 0
	for _, ch := range text {
		if byte(ch) == byte(a) {
			ac++
		} else if byte(ch) == byte(b) {
			bc++
			ans += ac
		}
	}
	return int64(ans + max(ac, bc))
}
