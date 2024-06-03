package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMaximumNumber(9, 1))
	fmt.Println(findMaximumNumber(7, 2))
	fmt.Println(findMaximumNumber(4096, 6))
	fmt.Println(findMaximumNumber(3278539330613, 5))
}

// 直接计算会超时
func findMaximumNumber(k int64, x int) int64 {
	ans := int64(0)
	v := 0

	for {
		nex := ans + 1
		c := cal(nex, x)
		if int64(v+c) > k {
			break
		}
		v = v + c
		ans = nex
	}
	return ans
}

func cal(n int64, k int) int {
	ans := 0
	i := 1
	for n>>(i*k-1) > 0 {
		if n>>(i*k-1)&1 == 1 {
			ans++
		}
		i++
	}
	return ans
}
