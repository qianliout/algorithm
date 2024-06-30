package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(calculateTax([][]int{{3, 50}, {7, 10}, {12, 25}}, 10))
}

func minPathCost(grid [][]int, moveCost [][]int) int {

	m, n := len(grid), len(grid[0])

	cost := make([][]int, m*n)
	for i := range cost {
		cost[i] = make([]int, n+1)
	}
	for i := range moveCost {
		for j, ch := range moveCost[i] {
			cost[i][j] = ch
		}
	}

	inf := math.MaxInt / 10
	var dfs func(i, j int) int
	mem := make([][]int, m+1)
	for i := range mem {
		mem[i] = make([]int, n+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	dfs = func(i, j int) int {
		// 终止条件
		if i == 0 {
			return grid[i][j]
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		va := grid[i][j]
		res := inf

		for k := 0; k < n; k++ {
			value := grid[i-1][k]
			res = min(res, dfs(i-1, k)+cost[value][j])
		}
		mem[i][j] = va + res
		return va + res
	}
	res := inf
	for j := 0; j < n; j++ {
		res = min(res, dfs(m-1, j))
	}
	return res
}

func calculateTax(brackets [][]int, income int) float64 {

	sort.Slice(brackets, func(i, j int) bool { return brackets[i][0] < brackets[j][0] })

	ans := 0
	flag := true
	for i := 0; i < len(brackets) && flag; i++ {
		up, per := brackets[i][0], brackets[i][1]
		prev := 0
		if i > 0 {
			prev = brackets[i-1][0]
		}
		cur := min(up, income) - prev
		ans += cur * per
		flag = up <= income
	}

	return float64(ans) / float64(100)
}
