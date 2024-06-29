package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(peopleAwareOfSecret(6, 2, 4))
	fmt.Println(peopleAwareOfSecret(4, 1, 3))
	fmt.Println(peopleAwareOfSecret(684, 18, 496))
}

func peopleAwareOfSecret(n int, delay int, forget int) int {
	mod := int(math.Pow10(9)) + 7
	cnt := make([]int, n)
	cnt[0] = 1
	for i := 0; i < n; i++ {
		for j := i + delay; j < min(n, i+forget); j++ {
			cnt[j] = (cnt[j] + cnt[i]) % mod
		}
	}
	res := 0
	for i := n - forget; i < n; i++ {
		res = (res + cnt[i]) % mod
	}

	return res % mod
}
