package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(minTimeToReach([][]int{{0, 4}, {4, 4}}))
	fmt.Println(minTimeToReach([][]int{{0, 0, 0}, {0, 0, 0}}))
}

func minTimeToReach(moveTime [][]int) int {
	m, n := len(moveTime), len(moveTime[0])
	dis := make([][]int, m)
	inf := 1 << 32
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = inf
		}
	}
	q := make(hp, 0)
	dis[0][0] = 0
	dirs := [][]int{{0, 1}, {0, -1}, []int{1, 0}, {-1, 0}}
	heap.Push(&q, item{i: 0, j: 0, dis: 0})
	for q.Len() > 0 {
		p := heap.Pop(&q).(item)
		i, j, d := p.i, p.j, p.dis
		if d > dis[i][j] {
			continue
		}
		// if i == m-1 && j == n-1 {
		// 	break
		// }
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if x < 0 || y < 0 || x >= m || y >= n {
				continue
			}
			// nd := d + moveTime[i][j]
			nd := max(d, moveTime[x][y]) + 1
			if dis[x][y] > nd {
				dis[x][y] = nd
				heap.Push(&q, item{i: x, j: y, dis: nd})
			}
		}
	}
	return dis[m-1][n-1]
}

type item struct {
	i, j, dis int
}

type hp []item

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(item)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
