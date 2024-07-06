package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSumDivThree([]int{1, 2, 3, 4, 4}))
}

func maxSumDivThree(nums []int) int {
	var dfs func(i, pre int) int
	n := len(nums)
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, 3)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	dfs = func(i, pre int) int {
		if i < 0 {
			if pre == 0 {
				return 0
			}
			return math.MinInt64
		}
		if mem[i][pre] != -1 {
			return mem[i][pre]
		}
		yes := dfs(i-1, (pre+nums[i])%3) + nums[i]
		no := dfs(i-1, pre)

		mem[i][pre] = max(yes, no)
		return mem[i][pre]
	}

	return dfs(n-1, 0)
}
