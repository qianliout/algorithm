package main

func main() {

}

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	g := make([][]int, n)
	for i, row := range graph {
		for _, j := range row {
			g[j] = append(g[j], i)
		}

	}
	var dfs func(i, fa int, path []int)
	ans := make([][]int, 0)
	dfs = func(i, fa int, path []int) {
		if i == 0 {
			path = append(path, i)
			ans = append(ans, reverse(path))
			return
		}
		path = append(path, i)
		for _, j := range g[i] {
			if j != fa {
				dfs(j, i, path)
			}
		}

		path = path[:len(path)-1]
	}
	dfs(n-1, -1, []int{})
	return ans
}

func reverse(path []int) []int {
	ans := make([]int, 0)
	for i := len(path) - 1; i >= 0; i-- {
		ans = append(ans, path[i])
	}
	return ans
}
