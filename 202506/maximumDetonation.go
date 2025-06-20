package main

func main() {

}

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	g := make([][]int, n) // g[i] 表示引爆i 这个炸弹时，还可以同时引爆那些炸弹
	for i, ch1 := range bombs {
		xi, yi, di := ch1[0], ch1[1], ch1[2]
		// g[i] = append(g[i], i)
		for j, ch2 := range bombs {
			if i == j {
				continue
			}
			xj, yj := ch2[0], ch2[1]
			if di*di >= (xi-xj)*(xi-xj)+(yi-yj)*(yi-yj) {
				g[i] = append(g[i], j)
			}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		visit := make([]bool, n)
		ans = max(ans, dfs(g, i, visit))
	}
	return ans
}

func dfs(g [][]int, x int, visit []bool) int {
	ans := 1
	visit[x] = true
	for _, y := range g[x] {
		if !visit[y] {
			ans += dfs(g, y, visit)
		}
	}
	return ans
}
