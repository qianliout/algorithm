package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(printVertically("TO BE OR NOT TO BE"))
}

func printVertically(s string) []string {
	split := strings.Split(s, " ")
	n := len(split)
	mx := 0
	for _, ch := range split {
		mx = max(mx, len(ch))
	}
	grid := make([][]byte, n)
	for i := range grid {
		grid[i] = make([]byte, mx)
		for j := range grid[i] {
			grid[i][j] = byte(' ')
		}
	}
	for i, word := range split {
		for j := 0; j < len(word); j++ {
			grid[i][j] = byte(word[j])
		}
	}
	ans := make([]string, 0)
	for j := 0; j < mx; j++ {
		str := make([]byte, 0)
		for i := 0; i < n; i++ {
			str = append(str, grid[i][j])
		}
		ans = append(ans, strings.TrimRight(string(str), " "))
	}
	return ans
}
