package main

import (
	"math"
)

func main() {

}

func minimumPossibleSum(n int, target int) int {
	m := min(target/2, n)
	l := m * (m + 1) / 2

	r := (target + (target + n - m - 1)) * (n - m) / 2
	mod := int(math.Pow10(9)) + 7
	return (l%mod + r%mod) % mod
}
