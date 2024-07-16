package main

import (
	"fmt"
)

func main() {
	// fmt.Println(numOfSubarrays([]int{1, 3, 5}))
	fmt.Println(numOfSubarrays([]int{1, 2, 3, 4, 5, 6, 7}))
}

// 结果不对
func numOfSubarrays(arr []int) int {
	n := 0
	for _, ch := range arr {
		n += ch & 1
	}
	ans := 0
	for i := 1; i <= n; i += 2 {
		ans += cal(n, i)
	}
	return ans
}

func cal(c, k int) int {
	return fac(c) / (fac(k) * fac(c-k))
}

func fac(n int) int {
	ans := 1
	for i := n; i > 0; i-- {
		ans = ans * i
	}
	return ans
}
