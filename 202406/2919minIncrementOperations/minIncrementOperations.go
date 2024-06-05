package main

import (
	"fmt"
)

func main() {
	fmt.Println(minIncrementOperations([]int{2, 3, 0, 0, 2}, 4))
	fmt.Println(minIncrementOperations([]int{0, 1, 3, 3}, 5))
	fmt.Println(minIncrementOperations([]int{4, 8, 20, 34, 21, 13, 45, 3, 42, 48, 49, 35, 48, 2, 11, 22, 33, 2, 32, 33, 50, 44, 40, 41, 27, 28, 46, 30, 43}, 86))
}

func minIncrementOperations(nums []int, k int) int64 {
	n := len(nums)
	mem := make([][]int64, n+1)
	for i := range mem {
		mem[i] = make([]int64, 4)
		for j := range mem[i] {
			mem[i][j] = int64(-n)
		}
	}
	var dfs func(i, j int) int64

	dfs = func(i, j int) int64 {
		if i < 0 {
			return 0
		}
		if mem[i][j] >= 0 {
			return mem[i][j]
		}
		// 选 i
		res := dfs(i-1, 0) + int64(max(0, k-nums[i]))
		// 不选
		if j < 2 {
			res = min(res, dfs(i-1, j+1))
		}
		mem[i][j] = res
		return res
	}
	return dfs(n-1, 0)
}

// func dfs(nums []int, k int, i, j int) int {
// 	if i < 0 {
// 		return 0
// 	}
// 	// 选 i
// 	res := dfs(nums, k, i-1, 0) + max(0, k-nums[i])
// 	// 不选
// 	if j < 2 {
// 		res = min(res, dfs(nums, k, i-1, j+1))
// 	}
// 	return res
// }
