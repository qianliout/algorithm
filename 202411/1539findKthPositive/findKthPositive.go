package main

import (
	"fmt"
)

func main() {
	// fmt.Println(findKthPositive([]int{2, 3, 4, 7, 11}, 5))
	fmt.Println(findKthPositive([]int{2}, 1))
}

func findKthPositive(arr []int, k int) int {
	n := len(arr)
	le, ri := 0, n
	for le < ri {
		mid := le + (ri-le)/2
		if mid >= 0 && mid <= n && arr[mid]-mid >= k+1 {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	return k + le
}

// 请你找到这个数组里第 k 个缺失的正整数
