package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numDifferentIntegers("00a123bc34d8ef34"))
}

func numDifferentIntegers(word string) int {
	ss := []byte(word)
	for i, ch := range word {
		if !(ch >= '0' && ch <= '9') {
			ss[i] = ' '
		}
	}
	cnt := make(map[string]int)
	split := strings.Split(string(ss), " ")
	for _, ch := range split {
		if ch == "" {
			continue
		}
		ch = strings.TrimLeft(ch, "0")
		if ch == "" {
			ch = "0"
		}
		cnt[ch]++
	}

	return len(cnt)
}
