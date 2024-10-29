package main

import (
	"fmt"
)

func main() {
	fmt.Println(passThePillow(9, 4))
}

func passThePillow(n int, t int) int {
	a, b := t/(n-1), t%(n-1)
	if a&1 == 0 {
		return b
	}
	return n - b
}
