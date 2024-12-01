package main

func main() {

}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {

	g := make([][]node, n)
	for _, e := range redEdges {
		g[e[0]] = append(g[e[0]], node{x: e[1], color: 0})
	}

	for _, e := range blueEdges {
		g[e[0]] = append(g[e[0]], node{x: e[1], color: 1})
	}

	queue := make([]node, 0)
	queue = append(queue, node{0, 1}, node{0, 0})
	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	visit := make([][2]bool, n)
	visit[0] = [2]bool{true, true} // 第0个表示红色，第1个表示blue
	level := 0
	for len(queue) > 0 {
		lev := make([]node, 0)
		for _, no := range queue {
			if dis[no.x] < 0 {
				dis[no.x] = level
			}
			for _, to := range g[no.x] {
				if to.color != no.color && !visit[to.x][to.color] {
					visit[to.x][to.color] = true
					lev = append(lev, to)
				}
			}
		}
		level++
		queue = lev
	}
	return dis
}

type node struct {
	x, color int
} // color: 0 表示路径最终到节点 x 的有向边为红色,1表示蓝色
