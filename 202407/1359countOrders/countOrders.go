package main

import (
	"math"
)

func main() {

}

func countOrders(n int) int {
	mod := int(math.Pow10(9)) + 7
	ans := 1
	for i := 2; i <= n; i++ {
		cur := i * (2*i - 1) % mod
		ans = (ans * cur) % mod
	}
	return ans % mod
}
