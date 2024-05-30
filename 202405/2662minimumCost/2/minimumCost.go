package main

import (
	"fmt"
	"math"
)

func main() {
	sp := [][]int{{1, 2, 3, 3, 2}, {3, 4, 4, 5, 1}}
	sp2 := [][]int{{3, 2, 3, 4, 4}, {3, 3, 5, 5, 5}, {3, 4, 5, 6, 6}}
	fmt.Println(minimumCost([]int{1, 1}, []int{4, 5}, sp))
	fmt.Println(minimumCost([]int{3, 2}, []int{5, 7}, sp2))
}

type pair struct {
	x, y int
}

func minimumCost(start []int, target []int, specialRoads [][]int) int {
	// 看做一个稠图
	inf := math.MaxInt
	visit := make(map[pair]bool)
	dis := make(map[pair]int, len(specialRoads)+2)
	fir := pair{start[0], start[1]}
	end := pair{target[0], target[1]}
	dis[end] = inf
	// fir 一定是最后赋值，因为 start 和end 可能是一个点
	dis[fir] = 0
	for {
		v, dv := pair{}, -1
		for p, d := range dis {
			if !visit[p] && (dv < 0 || d < dv) {
				v, dv = p, d
			}
		}

		if v == end { // 到终点的最短路已确定
			return dv
		}
		visit[v] = true
		dis[end] = min(dis[end], dv+end.x-v.x+end.y-v.y) // 更新到终点的最短路
		// 再更新这些特殊点
		for _, r := range specialRoads {
			w := pair{r[2], r[3]}
			d := dv + abs(r[0]-v.x) + abs(r[1]-v.y) + r[4]
			if dw, ok := dis[w]; !ok || d < dw {
				dis[w] = d
			}
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
