package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(putMarbles([]int{1, 3, 5, 1}, 2))
}

func putMarbles(wt []int, k int) int64 {
	for i := 0; i < len(wt)-1; i++ {
		wt[i] += wt[i+1]
	}
	wt = wt[:len(wt)-1]
	sort.Ints(wt)
	mi := 0
	for i := 0; i < k-1; i++ {
		mi += wt[i]
	}
	mx := 0
	for i := len(wt) - 1; i > len(wt)-k; i-- {
		mx += wt[i]
	}
	return int64(mx - mi)
}
