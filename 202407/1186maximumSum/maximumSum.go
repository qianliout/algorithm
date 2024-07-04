package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maximumSum([]int{1, -2, 0, 3}))
	fmt.Println(maximumSum([]int{1, -2, -2, 3}))
	fmt.Println(maximumSum([]int{-1, -1, -1, -1}))
	fmt.Println(maximumSum([]int{2, 1, -2, -5, -2}))
}

func maximumSum(nums []int) int {
	inf := math.MinInt / 2
	n := len(nums)
	var dfs func(i int, del int) int
	mem := make([][]int, n+1)
	for i := range mem {
		mem[i] = make([]int, 2)
		for j := range mem[i] {
			mem[i][j] = inf
		}
	}

	// del表示是否删除一个元素
	dfs = func(i int, del int) int {
		if i < 0 || i >= n {
			return inf
		}
		if mem[i][del] > inf {
			return mem[i][del]
		}

		// 一个元素都不删除
		if del == 0 {
			a := nums[i]               // 不选 i 左边的
			b := dfs(i-1, 0) + nums[i] // 选i 左边的
			mem[i][del] = max(a, b)

			return max(a, b)
		}
		// 需要删除一个元素
		a := dfs(i-1, 0)           // 删除 i,i 前面的就不能删除了
		b := dfs(i-1, 1) + nums[i] // 不删除 i,那么就只能在之前删除

		mem[i][del] = max(a, b)

		return max(a, b)
	}

	ans := inf
	for i := 0; i < n; i++ {
		ans = max(ans, max(dfs(i, 0), dfs(i, 1)))
	}
	return ans
}
