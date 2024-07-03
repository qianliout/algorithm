package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(movesToMakeZigzag([]int{1, 2, 3}))
	fmt.Println(movesToMakeZigzag([]int{9, 6, 1, 6, 2}))
}

func movesToMakeZigzag(nums []int) int {
	n := len(nums)
	ans1, ans2 := 0, 0
	arr := append([]int{math.MaxInt}, nums...)
	arr = append(arr, math.MaxInt)
	for i := 1; i <= n; i++ {
		ans := max(0, arr[i]-min(arr[i-1], arr[i+1])+1)
		if i&1 == 0 {
			ans1 += ans
		} else {
			ans2 += ans
		}
	}

	return min(ans1, ans2)
}
