package main

import (
	"fmt"
)

func main() {
	fmt.Println(pivotArray([]int{9, 5, 3, 10, 10, 14, 12}, 10))
}

func pivotArray(nums []int, pivot int) []int {
	pre, mid, suf := make([]int, 0), make([]int, 0), make([]int, 0)
	for _, ch := range nums {
		if ch == pivot {
			mid = append(mid, ch)
		} else if ch > pivot {
			suf = append(suf, ch)
		} else if ch < pivot {
			pre = append(pre, ch)
		}
	}
	pre = append(pre, mid...)
	pre = append(pre, suf...)
	return pre
}
