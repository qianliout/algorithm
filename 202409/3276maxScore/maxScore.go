package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(maxScore([][]int{{1, 2, 3}, {4, 3, 2}, {1, 1, 1}}))
}

func maxScore(grid [][]int) int {
	n := len(grid)
	// 因为题目中说了值最大是100,所以 max可以直接设置成100，而不用计算
	mx := grid[0][0]
	for _, v := range grid {
		for _, v2 := range v {
			mx = max(mx, v2)
		}
	}
	g := make(map[int][]int)
	for i, row := range grid {
		ro := slices.Compact(row)
		for _, r := range ro {
			g[r] = append(g[r], i)
		}
	}
	mem := make([][]int, mx+1)
	for j := range mem {
		mem[j] = make([]int, 1<<n)
		for k := range mem[j] {
			mem[j][k] = -1
		}
	}

	// i 表示要选的数，初始值就是最大值
	// 是表示选择的行集合,初始值就是0，因为都还没有选
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		// 数都选完了
		if i <= 0 {
			return 0
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		// 不选 i 这个数
		res := dfs(i-1, j)
		// 选i 这个数
		for _, v := range g[i] {
			if j&(1<<v) != 0 {
				continue
			}
			res = max(res, dfs(i-1, j|(1<<v))+i)
		}
		mem[i][j] = res
		return res
	}
	ans := dfs(mx, 0)
	return ans
}

// 1 <= grid.length, grid[i].length <= 10
// 1 <= grid[i][j] <= 100
