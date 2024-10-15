package main

import (
	"sort"
)

func main() {

}

func minimumTime(n int, edges [][]int, disappear []int) []int {
	g := make([][]pair, n)
	for _, ch := range edges {
		x, y, l := ch[0], ch[1], ch[2]
		g[x] = append(g[x], pair{to: y, le: l})
		g[y] = append(g[y], pair{to: x, le: l})
	}
	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	dis[0] = 0
	queue := []distance{{node: 0, dis: 0}}

	// 写一个朴素的得了，不用堆优化了
	// 对于本题目来说会超时
	for len(queue) > 0 {
		sort.Slice(queue, func(i, j int) bool { return queue[i].dis < queue[j].dis })
		qu := queue[0]
		no, di := qu.node, qu.dis
		queue = queue[1:]
		if dis[no] >= 0 && dis[no] < di {
			continue
		}
		for _, ch := range g[no] {
			y, wt := ch.to, ch.le
			nd := qu.dis + wt
			if nd < disappear[y] && (dis[y] < 0 || dis[y] > nd) {
				dis[y] = nd
				queue = append(queue, distance{node: y, dis: nd})
			}
		}
	}
	return dis
}

type pair struct {
	to, le int
}

type distance struct {
	node, dis int
}
