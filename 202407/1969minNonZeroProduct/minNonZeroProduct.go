package main

import (
	"fmt"
)

func main() {
	fmt.Println(minNonZeroProduct(3))

	fmt.Println(pow(10, 10), pow2(10, 10))

}

const mod = 1_000_000_007

func minNonZeroProduct(p int) int {
	k := 1<<p - 1
	return k % mod * pow2(k-1, p-1) % mod
}

func pow(x, p int) int {
	if p == 0 {
		return 1
	}
	// if p == 1 {
	// 	return x
	// }

	nex := pow(x, p/2)
	if p%2 == 0 {
		return nex * nex % mod
	}
	return x * nex * nex % mod
}

func pow2(x, p int) int {
	res := 1
	for x %= mod; p > 0; p-- {
		res = res * x % mod
		x = x * x % mod
	}
	return res
}
