package main

func main() {

}

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	ans := make([][]int, 0)
	dfs(graph, 0, n-1, []int{0}, &ans)
	return ans
}

func dfs(graph [][]int, start, end int, path []int, ans *[][]int) {
	if start == end {
		*ans = append(*ans, append([]int{}, path...))
		return
	}

	nex := graph[start]
	for _, ch := range nex {
		path = append(path, ch)
		dfs(graph, ch, end, path, ans)
		path = path[:len(path)-1]
	}
}
