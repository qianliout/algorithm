package main

import (
	"fmt"
)

func main() {
	fmt.Println(pivotInteger(8))
}

func pivotInteger(n int) int {
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + i
	}
	for i := 1; i <= n; i++ {
		if sum[i] == sum[n]-sum[i-1] {
			return i
		}
	}
	return -1
}
