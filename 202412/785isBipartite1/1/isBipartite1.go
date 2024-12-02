package main

func main() {

}

func isBipartite(graph [][]int) bool {
	n := len(graph)
	//  c 表示颜色，0表示还没有被染色，1表示红色，2 表示蓝色
	var dfs func(i, c int) bool

	g := make([]int, n)

	dfs = func(i, c int) bool {
		if g[i] == c {
			return true
		}
		nc := -c
		if g[i] == nc {
			return false
		}
		g[i] = c
		ans := true
		for _, k := range graph[i] {
			if g[k] == c {
				return false
			}
			if g[k] == nc {
				continue
			}
			ans = ans && dfs(k, nc)
		}
		return ans
	}
	ans := true
	for i := 0; i < n; i++ {
		g = make([]int, n)
		ans = ans && dfs(i, 1)
	}
	return ans
}

// 注意，如果一个图是二分图，那么从所有点出发，都能划分成二分图，而不只是从一个特定的点出发

// 不存在自环（graph[u] 不包含 u）。
// 不存在平行边（graph[u] 不包含重复值）。
// 如果 v 在 graph[u] 内，那么 u 也应该在 graph[v] 内（该图是无向图）
// 这个图可能不是连通图，也就是说两个节点 u 和 v 之间可能不存在一条连通彼此的路径。
