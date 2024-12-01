package main

func main() {

}

func findCircleNum(conn [][]int) int {
	n := len(conn)
	visit := make([]bool, n)
	var dfs func(i int)
	dfs = func(i int) {
		if i < 0 || i >= n {
			return
		}
		visit[i] = true
		for j := 0; j < n; j++ {
			if conn[i][j] == 1 && !visit[j] {
				dfs(j)
			}
		}
	}
	cnt := 0

	for i := 0; i < n; i++ {
		if !visit[i] {
			cnt++
			dfs(i)
		}
	}
	return cnt
}
