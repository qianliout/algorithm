package main

import (
	"fmt"
)

func main() {
	fmt.Println(numBusesToDestination([][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6))
}

func numBusesToDestination(routes [][]int, source int, target int) int {
	// 车站的公交车集合
	n := len(routes)

	station := make(map[int][]int, n)

	for i, route := range routes {
		for _, st := range route {
			station[st] = append(station[st], i)
		}
	}
	queue := make([]pire, 0)
	// visit1 表已示坐过的车
	// visit2 表已示已到达过的站
	visit1 := make(map[int]bool)
	visit2 := make(map[int]bool)

	// 先把终点加进去，从终点找起点
	queue = append(queue, pire{sta: source, cost: 0})
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		if first.sta == target {
			return first.cost
		}
		// 找这个站还没有坐过的车
		cars := station[first.sta]
		for _, c := range cars {
			if visit1[c] {
				continue
			}
			// 找这些车能到达的站
			visit1[c] = true
			for _, st := range routes[c] {
				if visit2[st] {
					continue
				}
				visit2[st] = true
				queue = append(queue, pire{sta: st, cost: first.cost + 1})
			}
		}
	}
	return -1
}

type pire struct {
	sta  int
	cost int
}

// 你之前可以坐一号线了，你后面兜兜转转(十号线->二号线->一号线)再上一号线是没有意义的(之前上一号线是比这样更优的)。
// 所以直接记录我们每次能坐的所有公交车，将它的所有站都标记为已到达(且加入队列)。
