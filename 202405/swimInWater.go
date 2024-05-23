package main

import (
	"fmt"
	"sort"

	. "outback/algorithm/common/unionfind"
)

func main() {
	fmt.Println(swimInWater([][]int{{0, 2}, {1, 3}}))
}

func swimInWater(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	uf := NewRankUnionFind(m*n + 1)

	steps := make([][]int, 0)

	for i := range grid {
		for j, ch := range grid[i] {
			pos := i*n + j
			if i+1 < m {
				w := max(ch, grid[i+1][j])
				steps = append(steps, []int{w, pos, pos + n})
			}
			if j+1 < n {
				w := max(grid[i][j+1], ch)
				steps = append(steps, []int{w, pos, pos + 1})
			}
		}
	}
	sort.Slice(steps, func(i, j int) bool { return steps[i][0] <= steps[j][0] })

	for _, ch := range steps {
		uf.Union(ch[1], ch[2])
		if uf.IsConnected(0, m*n-1) {
			return ch[0]
		}
	}
	return 0
}
