package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxSatisfaction([]int{-1, -8, 0, 5, -9}))
	fmt.Println(maxSatisfaction([]int{4, 3, 2}))
	fmt.Println(maxSatisfaction([]int{-1, -4, -5}))
}

func maxSatisfaction(s []int) int {
	sort.Slice(s, func(i, j int) bool { return s[i] > s[j] })
	f := 0
	n := len(s)
	sum := make([]int, n+1)
	for i, ch := range s {
		sum[i+1] = ch + sum[i]
	}
	for i := 1; i <= n; i++ {
		ch := sum[i]
		if ch <= 0 {
			break
		}
		f += ch
	}
	return f
}
