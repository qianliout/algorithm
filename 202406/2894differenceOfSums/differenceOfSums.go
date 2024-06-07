package main

import (
	"fmt"
)

func main() {
	fmt.Println(differenceOfSums(10, 3))
}

func differenceOfSums(n int, m int) int {
	a := (1 + n) * n
	b := (m + (n/m)*m) * (n / m)
	return (a - b - b) / 2
}
