package main

import (
	"fmt"
)

func main() {
	fmt.Println(minTimeToType("bza"))
}

func minTimeToType(word string) int {
	word = "a" + word
	cnt := 0
	n := len(word)

	for i := 1; i < n; i++ {
		a := abs(int(word[i]) - int(word[i-1]))
		cnt += min(a, 26-a)
		cnt++
	}
	return cnt
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
