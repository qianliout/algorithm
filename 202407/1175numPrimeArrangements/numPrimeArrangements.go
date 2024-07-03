package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numPrimeArrangements(100))
}

var prime []int
var N = 100

var Cnt map[int]int

func init() {
	Cnt = make(map[int]int)
	nums := make([]int, 101)
	for i := 2; i <= int(math.Sqrt(float64(N))); i++ {
		if nums[i] == -1 {
			continue
		}
		for j := 2; i*j <= N; j++ {
			nums[i*j] = -1
		}
	}
	for i := 2; i <= 100; i++ {
		if nums[i] != -1 {
			prime = append(prime, i)
		}
		Cnt[i] = len(prime)
	}
}

func numPrimeArrangements(n int) int {
	mod := int(math.Pow10(9)) + 7
	pri := Cnt[n]
	sub := n - pri
	ans := 1
	for i := pri; i > 0; i-- {
		ans = (ans * i) % mod
	}
	for i := sub; i > 0; i-- {
		ans = (ans * i) % mod
	}

	return ans % mod
}
