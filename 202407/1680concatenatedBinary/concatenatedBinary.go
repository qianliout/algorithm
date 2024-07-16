package main

import (
	"fmt"
	"math"
	"math/bits"
)

func main() {
	fmt.Println(concatenatedBinary(3))
	fmt.Println(concatenatedBinary(12))
}

var mod = int(math.Pow10(9)) + 7

func concatenatedBinary(n int) int {
	res := dfs(n)
	return res
}

func dfs(x int) int {
	if x == 1 {
		return 1
	}
	pre := dfs(x - 1)
	l := bits.Len(uint(x))
	return (pre<<l + x) % mod
}
