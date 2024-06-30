package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumObstacles([][]int{{0, 1, 1}, {1, 1, 0}, {1, 1, 0}}))
	fmt.Println(minimumObstacles([][]int{{0, 1, 0, 0, 0}, {0, 1, 0, 1, 0}, {0, 0, 0, 1, 0}}))
}

/*
提示 1
把障碍物当作可以经过的单元格，经过它的代价为 1，空单元格经过的代价为 0。
提示 2
问题转化成从起点到终点的最短路。
*/
func minimumObstacles(grid [][]int) int {
	inf := math.MaxInt / 100
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = inf
		}
	}
	dis[0][0] = grid[0][0]
	queue := make(MinHeap, 0)
	queue = append(queue, pair{X: 0, Y: 0, Dis: grid[0][0]})
	for len(queue) > 0 {
		pop := heap.Pop(&queue).(pair)
		x, y := pop.X, pop.Y
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if !in(m, n, nx, ny) {
				continue
			}
			g := grid[nx][ny]
			if dis[x][y]+g < dis[nx][ny] {
				dis[nx][ny] = dis[x][y] + g
				heap.Push(&queue, pair{
					X:   nx,
					Y:   ny,
					Dis: dis[nx][ny],
				})
			}
		}
	}

	return dis[m-1][n-1]
}

func in(m, n, c, r int) bool {
	if c < 0 || r < 0 {
		return false
	}
	if c >= m || r >= n {
		return false
	}
	return true
}

type pair struct {
	X, Y, Dis int
}

type MinHeap []pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Dis <= h[j].Dis }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
