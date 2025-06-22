package main

func main() {

}

func allPathsSourceTarget(graph [][]int) [][]int {
	//

	n := len(graph)
	ans := make([][]int, 0)
	var dfs func(node int, path []int)
	dfs = func(node int, path []int) {
		path = append(path, node)
		if node == n-1 {
			ans = append(ans, Copy(path))
			return
		}
		for _, j := range graph[node] {
			dfs(j, path)
		}
		// 这里写不写是一样的，都可以得到正确的答案
		// path = path[:len(path)-1]
	}
	dfs(0, []int{})
	return ans
}

func Copy(path []int) []int {
	res := make([]int, len(path))
	for i := range path {
		res[i] = path[i]
	}
	return res
}
