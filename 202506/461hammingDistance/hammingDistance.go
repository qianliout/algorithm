package main

import (
	"fmt"
)

func main() {
	fmt.Println(hammingDistance(1, 4))
}

func hammingDistance(x int, y int) int {
	ans := 0
	for x != 0 || y != 0 {
		if x&1 != y&1 {
			ans++
		}
		x = x >> 1
		y = y >> 1
	}
	return ans
}
