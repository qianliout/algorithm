package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumTime([][]int{{0, 1, 3, 2}, {5, 1, 2, 5}, {4, 3, 8, 6}}))
}

func minimumTime(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	if grid[0][1] > 1 && grid[1][0] > 1 {
		return -1
	}
	ans := dijkstra(grid)
	return ans
}

func dijkstra(grid [][]int) int {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	m, n := len(grid), len(grid[0])
	inf := math.MaxInt / 10
	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = inf
		}
	}
	dis[0][0] = 0
	hp := make(MinHeap, 0)
	heap.Push(&hp, Item{X: 0, Y: 0, D: 0})

	for len(hp) > 0 {
		pop := heap.Pop(&hp).(Item)
		x, y, d := pop.X, pop.Y, pop.D
		// 这个判断可有可没有
		// if d > dis[x][y] {
		// 	continue
		// }

		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx < 0 || ny < 0 || nx >= m || ny >= n {
				continue
			}
			nd := max(d+1, grid[nx][ny])
			nd += (nd - nx - ny) % 2 // 可以通过反复横跳来等待时间，这样的话，nd 必须和 nx+ny 同奇偶
			if nd < dis[nx][ny] {
				dis[nx][ny] = nd
				heap.Push(&hp, Item{X: nx, Y: ny, D: nd})
			}
		}
	}

	return dis[m-1][n-1]
}

type Item struct {
	X, Y int
	D    int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].D < h[j].D }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
