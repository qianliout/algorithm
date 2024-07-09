package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMinFibonacciNumbers(19))
}

func findMinFibonacciNumbers(k int) int {
	a, b := 1, 1
	for b < k {
		b, a = a+b, b
	}
	ans := 0
	for k != 0 {
		if k >= b {
			k -= b
			ans++
		}
		b, a = a, b-a
	}
	return ans
}
