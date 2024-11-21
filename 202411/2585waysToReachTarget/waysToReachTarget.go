package main

import (
	"math"
)

func main() {

}

func waysToReachTarget(target int, types [][]int) int {
	mod := int(math.Pow10(9)) + 7
	ans := 0

	return ans % mod
}

type pair struct {
	count int
	marks int
}
