package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestDecomposition("ghiabcdefhelloadamhelloabcdefghi"))
}

func longestDecomposition(text string) int {
	return dfs(text)
}

func dfs(text string) int {
	if text == "" {
		return 0
	}
	start := 0
	end := len(text) - 1
	for start < end {
		if text[:start+1] == text[end:] {
			return 2 + dfs(text[start+1:end])
		}
		start++
		end--
	}
	return 1
}
