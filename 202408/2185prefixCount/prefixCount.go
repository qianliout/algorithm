package main

import (
	"strings"
)

func main() {

}

func prefixCount(words []string, pref string) int {
	ans := 0
	for _, ch := range words {
		//if ch != pref && strings.HasPrefix(ch, pref) {
		if strings.HasPrefix(ch, pref) {
			ans++
		}
	}
	return ans
}
