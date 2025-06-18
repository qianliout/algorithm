package main

func findShortestCycle(n int, edges [][]int) int {
	g := make([][]int, n)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	inf := 1 << 32
	bfs := func(start int) int {
		//  每次bfs 都要初始化
		dis := make([]int, n)
		for i := range dis {
			dis[i] = -1
		}
		ans := inf
		dis[start] = 0
		q := []node{{idx: start, fa: -1}}
		for len(q) > 0 {
			fir := q[0]
			q = q[1:]
			for _, y := range g[fir.idx] {
				if dis[y] < 0 {
					dis[y] = dis[fir.idx] + 1
					q = append(q, node{idx: y, fa: fir.idx})
				} else if y != fir.fa {
					// 说明是环
					ans = min(ans, dis[fir.idx]+dis[y]+1)
				}
			}
		}
		return ans
	}
	ans := inf
	for i := 0; i < n; i++ {
		ans = min(ans, bfs(i))
	}
	if ans < inf {
		return ans
	}
	return -1
}

type node struct {
	idx int
	fa  int
}
