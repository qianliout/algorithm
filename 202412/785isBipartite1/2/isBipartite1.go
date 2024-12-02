package main

func main() {

}

func isBipartite(graph [][]int) bool {
	n := len(graph)

	visit := make([]int, n) // 0 未染色，1：红色，2：蓝色
	q := make([]int, 0)
	for i := 0; i < n; i++ {
		// 说明已染色
		if visit[i] != 0 {
			continue
		}
		visit[i] = 1
		q = append(q, i)

		for len(q) > 0 {
			fir := q[0]
			q = q[1:]
			c := visit[fir]
			nc := -c
			for _, j := range graph[fir] {
				if visit[j] == c {
					return false
				}
				if visit[j] == 0 {
					visit[j] = nc
					q = append(q, j)
				}
			}
		}
	}
	return true
}

// 注意，如果一个图是二分图，那么从所有点出发，都能划分成二分图，而不只是从一个特定的点出发

// 不存在自环（graph[u] 不包含 u）。
// 不存在平行边（graph[u] 不包含重复值）。
// 如果 v 在 graph[u] 内，那么 u 也应该在 graph[v] 内（该图是无向图）
// 这个图可能不是连通图，也就是说两个节点 u 和 v 之间可能不存在一条连通彼此的路径。
