package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	// fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
	fmt.Println(mostCommonWord("a, a, a,c,c,c,, a, b,,,b,b,c, c,,,", []string{"b"}))
}

func mostCommonWord(paragraph string, banned []string) string {
	bb := make(map[string]bool)
	for _, w := range banned {
		w = strings.ToLower(w)
		bb[w] = true
	}
	words := split(paragraph)
	mc, mw := 0, ""
	wordMap := make(map[string]int)
	for _, w := range words {

		if bb[w] {
			continue
		}
		wordMap[w]++
		if wordMap[w] > mc {
			mc = wordMap[w]
			mw = w
		}
	}
	return mw
}

func split(s string) []string {
	sym := []string{" ", "!", "?", "'", ",", ";", "."}
	ss := []byte(s)
	start := 0
	ans := make([]string, 0)

	for i := 0; i < len(s); i++ {
		if slices.Contains(sym, string(ss[i])) {
			w := strings.ToLower(string(ss[start:i]))
			if w != "" {
				ans = append(ans, w)
			}
			start = i + 1
		}
	}
	if start <= len(ss) {
		w := strings.ToLower(string(ss[start:]))
		if w != "" {
			ans = append(ans, w)
		}
	}

	return ans
}
