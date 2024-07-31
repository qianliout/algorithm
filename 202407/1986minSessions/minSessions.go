package main

import (
	"fmt"
)

func main() {
	fmt.Println(minSessions([]int{3, 1, 3, 1, 1}, 8))
	fmt.Println(minSessions([]int{9, 8, 8, 4, 6}, 14))
}

func minSessions(tasks []int, sessionTime int) int {
	n := len(tasks)
	m := 1 << n
	sum := make([]int, m)
	for i, ch := range tasks {
		k := 1 << i
		for j := 0; j < k; j++ {
			sum[j|k] = sum[j] + ch
		}
	}

	f := make([]int, m)
	for i := 0; i < m; i++ {
		f[i] = n
	}
	f[0] = 0
	for i := 0; i < m; i++ {
		for sub := i; sub > 0; sub = (sub - 1) & i {
			if sum[sub] <= sessionTime {
				f[i] = min(f[i], f[i^sub]+1)
			}
		}
	}
	return f[m-1]
}
