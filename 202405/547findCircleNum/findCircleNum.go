package main

func main() {

}

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	vis := make([]bool, n)
	cnt := 0
	for i := 0; i < n; i++ {
		if vis[i] {
			continue
		}
		cnt++
		dfs(isConnected, n, i, vis)
	}

	return cnt
}

func dfs(c [][]int, cites int, start int, vis []bool) {
	for j := 0; j < cites; j++ {
		if vis[j] {
			continue
		}
		if c[start][j] == 0 {
			continue
		}
		vis[j] = true
		dfs(c, cites, j, vis)
	}
}
