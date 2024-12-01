package main

func main() {

}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {

	g := make([][]node, n)
	for _, ch := range redEdges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], node{idx: y, color: "red"})
	}

	for _, ch := range blueEdges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], node{idx: y, color: "blue"})
	}

	queue := make([]node, 0)
	queue = append(queue, node{0, "red"}, node{0, "blue"})
	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	redVisit := make([]bool, n)
	blueVisit := make([]bool, n)

	redVisit[0] = true
	blueVisit[0] = true

	level := 0

	for len(queue) > 0 {
		lev := make([]node, 0)
		for _, no := range queue {
			if dis[no.idx] < 0 {
				dis[no.idx] = level
			}
			for _, to := range g[no.idx] {
				if to.color != no.color {
					if to.color == "red" && !redVisit[to.idx] {
						redVisit[to.idx] = true
						lev = append(lev, to)
					}
					if to.color == "blue" && !blueVisit[to.idx] {
						blueVisit[to.idx] = true
						lev = append(lev, to)
					}
				}
			}
		}
		level++
		queue = lev
	}
	return dis
}

type node struct {
	idx   int    // idx
	color string // 出站的颜色
}
