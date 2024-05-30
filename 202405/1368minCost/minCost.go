package main

import (
	"container/heap"
	"fmt"
	"math"

	. "outback/algorithm/common/commonHeap"
)

func main() {
	fmt.Println(minCost([][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {1, 1, 1, 1}, {2, 2, 2, 2}}))
}

// 求翻转最少，不就是找一条路径，边权最小吗
func minCost(grid [][]int) int {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	inf := math.MaxInt32 / 10
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	col, row := len(grid), len(grid[0])
	n := col * row

	// 一个点最多和四个点相边，所以不初始化成[][]的格式可以节约内存，也减少相邻边的计算
	g := make([]map[int]int, n)
	for i := range g {
		g[i] = make(map[int]int)
	}
	// 初始边权
	for i := range grid {
		for j, ch := range grid[i] {
			for k, dir := range dirs {
				// 一个点最多初始化4点相临的点，这些点的边权要么是0，要么是1
				nx, ny := i+dir[0], j+dir[1]
				if nx < 0 || nx >= col || ny < 0 || ny >= row {
					continue
				}
				nex := pos(nx, ny, row)
				if k+1 == ch {
					g[pos(i, j, row)][nex] = 0
				} else {
					g[pos(i, j, row)][nex] = 1
				}
			}
		}
	}
	hp := make(PriorityQueue, 0)
	heap.Push(&hp, &IntItem{Value: 0, Priority: 0})
	visit := make([]bool, n)
	dis := make([]int, n)
	for i := range dis {
		dis[i] = inf
	}
	dis[0] = 0
	for hp.Len() > 0 {
		to := heap.Pop(&hp).(*IntItem)
		x := to.Value
		if visit[x] {
			continue
		}
		visit[x] = true
		for y, d := range g[x] {
			if y == x {
				continue
			}
			dis[y] = min(dis[y], -to.Priority+d)
			heap.Push(&hp, &IntItem{Value: y, Priority: -dis[y]})
		}
	}
	return dis[n-1]
}

func pos(c, r, row int) int {
	return c*row + r
}
