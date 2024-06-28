package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSumAfterPartitioning([]int{1, 15, 7, 9, 2, 5, 10}, 3))
	fmt.Println(maxSumAfterPartitioning([]int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}, 4))
}

func maxSumAfterPartitioning(arr []int, k int) int {
	var dfs func(i int) int
	n := len(arr)
	mem := make([]int, n)
	for i := range mem {
		mem[i] = -1
	}

	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if mem[i] >= 0 {
			return mem[i]
		}
		res := 0
		mx := 0
		for j := 0; j < k && i-j >= 0; j++ {
			mx = max(mx, arr[i-j])
			res = max(res, dfs(i-j-1)+mx*(j+1))
		}
		mem[i] = res
		return res
	}
	return dfs(n - 1)
}
