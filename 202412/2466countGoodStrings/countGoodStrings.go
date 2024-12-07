package main

import (
	"math"
)

func main() {

}
func countGoodStrings(low int, high int, zero int, one int) int {
	f := make([]int, high+1)
	f[0] = 1
	for i := 1; i <= high; i++ {
		if i >= zero {
			f[i] += f[i-zero]
		}
		if i >= one {
			f[i] += f[i-one]
		}
	}
	mod := int(math.Pow10(9)) + 7
	ans := 0
	for i := low; i <= high; i++ {
		ans = (ans + f[i]) % mod
	}
	return ans
}
