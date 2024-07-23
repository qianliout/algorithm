package main

import (
	"fmt"
)

func main() {
	fmt.Println(reinitializePermutation(6))
}

// 暴力的做法
func reinitializePermutation(n int) int {
	perm := make([]int, n)
	for i := 0; i < n; i++ {
		perm[i] = i
	}
	step := 0
	for {
		arr := make([]int, n)
		for j := 0; j < n; j++ {
			if j%2 == 0 {
				arr[j] = perm[j/2]
			} else {
				arr[j] = perm[n/2+(j-1)/2]
			}
		}
		step++
		k := 0
		for ; k < n; k++ {
			if arr[k] != k {
				break
			}
		}
		if k == n {
			return step
		}
		for i := 0; i < n; i++ {
			perm[i] = arr[i]
		}
	}
}
