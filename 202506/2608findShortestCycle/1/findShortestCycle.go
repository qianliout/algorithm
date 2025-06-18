package main

func findShortestCycle(n int, edges [][]int) int {
	g := make([][]int, n)

	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	ans := 1 << 32
	for i := 0; i < n; i++ {
		ans = min(ans, bfs(g, n, i))
	}
	if ans >= 1<<32 {
		return -1
	}
	return ans
}

type node struct {
	x  int
	fa int
}

func bfs(g [][]int, n, start int) int {
	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	dis[start] = 0
	q := []node{{x: start, fa: -1}}
	ans := 1 << 32

	for len(q) > 0 {
		fir := q[0]
		q = q[1:]

		for _, y := range g[fir.x] {
			if dis[y] < 0 {
				// 说明第一次访问
				dis[y] = dis[fir.x] + 1
				q = append(q, node{x: y, fa: fir.x})
			} else if y != fir.fa {
				// 不是父节点,且已访问过，说明有环
				// +1 是因为x到y的这个边
				ans = min(ans, dis[fir.x]+dis[y]+1)
			}
		}
	}
	return ans
}
