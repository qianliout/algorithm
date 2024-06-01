package main

import (
	"fmt"
)

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}

func pivotIndex(nums []int) int {
	pre := make([]int, len(nums)+1)
	sum := 0
	for i, ch := range nums {
		pre[i+1] = pre[i] + ch
		sum += ch
	}
	fmt.Println(pre, sum)
	for i := 0; i < len(pre)-1; i++ {
		if pre[i] == sum-pre[i+1] {
			return i
		}
	}
	return -1
}
