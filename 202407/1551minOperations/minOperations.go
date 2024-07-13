package main

import (
	"fmt"
)

func main() {
	fmt.Println(minOperations(3))
	fmt.Println(minOperations(6))
}

func minOperations(n int) int {
	op := 0
	for i := 1; i <= n; i += 2 {
		op += n - i
	}
	return op
}
