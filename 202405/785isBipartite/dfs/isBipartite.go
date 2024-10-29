package main

func main() {

}

func isBipartite1(graph [][]int) bool {
	n := len(graph)
	visit := make([]int, n) // 0 未染色，1蓝色 -1黄色

	var dfs func(x, c int) bool
	visit = make([]int, n)
	dfs = func(x, c int) bool {
		visit[x] = c
		for _, y := range graph[x] {
			if visit[y] == c {
				return false
			}
			if visit[y] == 0 && !dfs(y, -c) {
				return false
			}
		}
		return true
	}

	for i := 0; i < n; i++ {
		// 每次都新开一个 visit 数组，效率不高，可以考虑使用时间戳的方式
		visit = make([]int, n)
		if !dfs(i, 1) {
			return false
		}
	}
	return true
}

func isBipartite(graph [][]int) bool {
	n := len(graph)
	visit := make([]int, n) // 0 未染色，1蓝色 -1黄色

	var dfs func(x, c int) bool
	visit = make([]int, n)
	dfs = func(x, c int) bool {
		visit[x] = c
		for _, y := range graph[x] {
			if visit[y] == c {
				return false
			}
			// visit[y]!=abs(c) // 说明不是本次 dfs 的颜色，所以可以染色
			if visit[y] != abs(c) && !dfs(y, -c) {
				return false
			}
		}
		return true
	}
	ti := 0
	for i := 0; i < n; i++ {
		ti++
		if !dfs(i, ti) {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
