package main

import (
	"fmt"
)

func main() {
	fmt.Println(stoneGameV([]int{6, 2, 3, 4, 5, 5}))
	fmt.Println(stoneGameV([]int{1, 1, 2}))
}

func stoneGameV(stoneValue []int) int {
	var dfs func(le, ri int) int
	n := len(stoneValue)
	sum := make([]int, n+1)
	mem := make([][]int, n+1)
	for i := range mem {
		mem[i] = make([]int, n+1)
		for j := range mem {
			mem[i][j] = -1
		}
	}

	for i, ch := range stoneValue {
		sum[i+1] = sum[i] + ch
	}
	dfs = func(le, ri int) int {
		if le >= ri {
			return 0
		}
		if mem[le][ri] != -1 {
			return mem[le][ri]
		}
		ret := 0
		for i := le; i < ri; i++ {
			pre1 := sum[i+1] - sum[le]
			pre2 := sum[ri+1] - sum[i+1]
			if min(pre1, pre2)*2 < ret {
				break // 剪枝，不写这里也能过
			}

			if pre1 > pre2 {
				ret = max(ret, pre2+dfs(i+1, ri))
			} else if pre1 < pre2 {
				ret = max(ret, pre1+dfs(le, i))
			} else if pre1 == pre2 { // 这一步最容易出错
				ret = max(ret, pre2+dfs(i+1, ri), pre1+dfs(le, i))
			}
		}
		mem[le][ri] = ret
		return ret
	}
	return dfs(0, n-1)
}
