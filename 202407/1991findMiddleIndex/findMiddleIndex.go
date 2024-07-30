package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMiddleIndex([]int{2, 3, -1, 8, 4}))
	fmt.Println(findMiddleIndex([]int{2, 5}))
	fmt.Println(findMiddleIndex([]int{5}))
}

func findMiddleIndex(nums []int) int {
	n := len(nums)
	sum := make([]int, n+1)
	for i, ch := range nums {
		sum[i+1] = sum[i] + ch
	}

	for i := 0; i < n; i++ {
		le := sum[i]
		ri := sum[n] - sum[i+1]
		if le == ri {
			return i
		}
	}
	return -1
}
