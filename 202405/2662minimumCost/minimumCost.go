package main

import (
	"fmt"
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
	// 建图,使用邻接表
	g := make(map[pair]map[pair]int)
	st := pair{start[0], start[1]}
	end := pair{target[0], target[1]}
	visit := make(map[pair]bool)
	visit[st] = false
	visit[end] = false
	g[st] = make(map[pair]int)
	g[end] = make(map[pair]int)

	for _, ch := range specialRoads {
		sp1 := pair{ch[0], ch[1]}
		sp2 := pair{ch[2], ch[3]}
		visit[sp1] = false
		visit[sp2] = false
		if g[sp1] == nil {
			g[sp1] = make(map[pair]int)
		}
		if g[sp2] == nil {
			g[sp2] = make(map[pair]int)
		}
		g[sp1][sp2] = ch[4] // sp1--sp2
		g[st][sp1] = mn(st, sp1)
		g[st][sp2] = mn(st, sp2)
	}

	dis := make(map[pair]int)
	dis[st] = 0

	for {
		x := pair{x: -1}
		for i, ok := range visit {
			if ok {
				continue
			}
			d, ok2 := dis[i]
			if !ok2 && d < dis[x] {
				x = i
			}
		}
		// 返问完了
		if x.x <= -1 {
			break
		}
		visit[x] = true
		// 求值
		for nex, d := range g[x] {
			if _, ok := dis[nex]; !ok {
				dis[nex] = dis[x] + d
			} else {
				dis[nex] = min(dis[nex], dis[x]+d)
			}
		}
	}
	return dis[end]
}

func mn(st pair, ed pair) int {
	return abs(st.x-ed.x) + abs(st.y-ed.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
