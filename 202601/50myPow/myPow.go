package main

import (
	"fmt"
)

func main() {
	fmt.Println(-3 % 2)
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	mid, left := n/2, n%2
	c := myPow(x, mid)
	if left == 1 {
		return c * c * x
	}
	return c * c
}
