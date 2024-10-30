package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(primeSubOperation([]int{998, 2}))
	fmt.Println(primeSubOperation([]int{5, 8, 3}))
}

func primeSubOperation(nums []int) bool {
	// 1 <= nums[i] <= 1000
	prim := getPrime(1000)
	pre := 0
	for _, ch := range nums {
		if ch <= pre {
			return false
		}
		k := sort.SearchInts(prim, ch-pre)
		pre = ch - prim[k-1]
	}
	return true
}

func getPrime(n int) []int {
	pri := make([]bool, n)
	for i := range pri {
		pri[i] = true
	}
	pri[0], pri[1] = false, false
	for d := 2; d*d < n; d++ {
		for i := d * d; i < n; i += d {
			pri[i] = false
		}
	}
	ans := []int{0}
	for i, ch := range pri {
		if ch {
			ans = append(ans, i)
		}
	}
	return ans
}
