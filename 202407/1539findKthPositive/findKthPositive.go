package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(findKthPositive([]int{2, 3, 4, 7, 11}, 5))
}

func findKthPositive(arr []int, k int) int {
	start := 0
	cnt := 0
	n := len(arr)

	for i := 0; i < n; i++ {
		if arr[i] > start {
			for j := start + 1; j < arr[i]; j++ {
				cnt++
				if cnt >= k {
					return j
				}
			}
		}
		start = arr[i]
	}
	return slices.Max(arr) + k - cnt
}
