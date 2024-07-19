package main

import (
	"fmt"
)

func main() {
	fmt.Println(distinctPrimeFactors([]int{2, 4, 3, 7, 10, 6}))
}

func distinctPrimeFactors(nums []int) int {
	used := make(map[int]int)
	for _, ch := range nums {
		ans := find(ch)
		for _, j := range ans {
			used[j]++
		}
	}
	return len(used)
}

func find(n int) []int {
	ans := make([]int, 0)
	x := n
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			ans = append(ans, i)
		}
		// 然后把这个 i 能整除就全部整除完
		// for ; x%i == 0; x = x / i {
		// }
		for x%i == 0 {
			x = x / i
		}
	}
	if x > 1 {
		ans = append(ans, x)
	}
	return ans
}
