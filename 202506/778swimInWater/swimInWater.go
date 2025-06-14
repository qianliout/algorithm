package main

import (
	"sort"
)

func main() {

}

// 在一个 n x n 的整数矩阵 grid 中，每一个方格的值 grid[i][j] 表示位置 (i, j) 的平台高度。
// 当开始下雨时，在时间为 t 时，水池中的水位为 t 。你可以从一个平台游向四周相邻的任意一个平台，但是前提是此时水位必须同时淹没这两个平台。假定你可以瞬间移动无限距离，也就是默认在方格内部游动是不耗时的。当然，在你游泳的时候你必须待在坐标方格里面。
// 你从坐标方格的左上平台 (0，0) 出发。返回 你到达坐标方格的右下平台 (n-1, n-1) 所需的最少时间 。

func swimInWater(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	uf := NewSizeUnionFind(m*n + 1)

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
		if uf.IsCollect(0, m*n-1) {
			return ch[0]
		}
	}
	return 0
}

type UnionFind struct {
	Parent []int // 可以理解成下标i的最终父节点就是 Parent[i]
	Count  int
}

func NewSizeUnionFind(totalNodes int) *UnionFind {
	p := make([]int, totalNodes)
	for i := 0; i < totalNodes; i++ {
		p[i] = i
	}
	return &UnionFind{Parent: p, Count: totalNodes}
}

func (uf *UnionFind) Find(x int) int {
	if uf.Parent[x] != x {
		uf.Parent[x] = uf.Find(uf.Parent[x]) // 路径压缩
	}
	return uf.Parent[x]
}

func (uf *UnionFind) IsCollect(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func (uf *UnionFind) Union(x, y int) {
	xRoot := uf.Find(x)
	yRoot := uf.Find(y)

	if xRoot != yRoot {
		uf.Count--
		uf.Parent[yRoot] = xRoot // 直接把y的根节点指向x的根节点
	}
}
