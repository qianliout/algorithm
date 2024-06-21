package main

import (
	"math"
)

func main() {

}

// 不加 mem 会超时
func stoneGame(nums []int) bool {
	n := len(nums)
	inf := math.MinInt / 2
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, n)
		for j := range mem[i] {
			mem[i][j] = inf
		}
	}

	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		if i == j {
			return nums[i]
		}
		if mem[i][j] != inf {
			return mem[i][j]
		}
		left := nums[i] - dfs(i+1, j)
		right := nums[j] - dfs(i, j-1)

		a := max(left, right)
		mem[i][j] = a
		return a
	}
	ans := dfs(0, len(nums)-1)
	return ans > 0
}
