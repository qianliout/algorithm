package main

import (
	"fmt"
)

func main() {
	fmt.Println(waysToSplitArray([]int{10, 4, -8, 7}))
}

func waysToSplitArray(nums []int) int {
	n := len(nums)
	sum := make([]int, n+1)
	for i, ch := range nums {
		sum[i+1] = sum[i] + ch
	}
	ans := 0
	for i := 0; i < n-1; i++ {
		if sum[i+1] >= sum[n]-sum[i+1] {
			ans++
		}
	}
	return ans
}
